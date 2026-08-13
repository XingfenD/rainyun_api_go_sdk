// Package ssl implements the ry ssl command.
package ssl

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/XingfenD/rainyun_api_go_sdk/apis/common"
	apisssl "github.com/XingfenD/rainyun_api_go_sdk/apis/ssl"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/cliutil"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/model"
	"github.com/XingfenD/rainyun_api_go_sdk/cmd/ry/internal/output"
	"github.com/XingfenD/rainyun_api_go_sdk/sdk"

	"github.com/spf13/cobra"
)

func Cmd(rySDK **sdk.RainyunSDK, out **output.Printer) *cobra.Command {
	sslCmd := &cobra.Command{Use: "ssl", Short: "Manage Rainyun SSL certificates"}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List SSL certificates",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetSSLCertificateList(&apisssl.GetSslCertificateListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			items := make([]model.SslCertRecord, len(resp.Data.Records))
			for i, r := range resp.Data.Records {
				items[i] = toSslCertRecord(r)
			}
			return (*out).Print(items)
		},
	}

	getCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get SSL certificate details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetSslDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toSslCertDetail(resp.Data))
		},
	}

	uploadCmd := &cobra.Command{
		Use:   "upload <cert-file> <key-file>",
		Short: "Upload an SSL certificate",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cert, err := os.ReadFile(args[0])
			if err != nil {
				return err
			}
			key, err := os.ReadFile(args[1])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).UploadSSLCertificate(string(cert), string(key)); err != nil {
				return err
			}
			fmt.Println("SSL certificate uploading")
			return nil
		},
	}

	replaceCmd := &cobra.Command{
		Use:   "replace <id> <cert-file> <key-file>",
		Short: "Replace an SSL certificate",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			cert, err := os.ReadFile(args[1])
			if err != nil {
				return err
			}
			key, err := os.ReadFile(args[2])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).ReplaceSsl(id, string(cert), string(key)); err != nil {
				return err
			}
			fmt.Printf("SSL certificate %s replaced\n", args[0])
			return nil
		},
	}

	deleteCmd := &cobra.Command{
		Use:   "delete <id>",
		Short: "Delete an SSL certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).DeleteSsl(id); err != nil {
				return err
			}
			fmt.Printf("SSL certificate %s deleted\n", args[0])
			return nil
		},
	}

	applyCmd := &cobra.Command{Use: "apply", Short: "Manage free SSL certificate applications"}

	applyListCmd := &cobra.Command{
		Use:   "list",
		Short: "List SSL certificate applications",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetSSLCertApplyList(&apisssl.GetSSLOrderListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			return (*out).Print(model.RawData{Data: resp.Data})
		},
	}

	applyCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Apply for a free SSL certificate",
		RunE: func(cmd *cobra.Command, args []string) error {
			domains, _ := cmd.Flags().GetString("domains")
			verifyMethod, _ := cmd.Flags().GetString("verify-method")
			if domains == "" || verifyMethod == "" {
				return fmt.Errorf("--domains and --verify-method flags are required")
			}
			if _, err := (*rySDK).ApplyFreeSSLCertificate(&apisssl.ApplyFreeSSLCertRequest{
				Domains: domains, VerifyMethod: verifyMethod,
			}); err != nil {
				return err
			}
			fmt.Println("SSL certificate applying")
			return nil
		},
	}

	applyVerifyCmd := &cobra.Command{
		Use:   "verify <order-id>",
		Short: "Verify a free SSL certificate application",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			orderID, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).VerifyFreeSSLCertificate(orderID); err != nil {
				return err
			}
			fmt.Printf("Application %s verified\n", args[0])
			return nil
		},
	}

	orderCmd := &cobra.Command{Use: "order", Short: "Manage SSL certificate orders"}

	orderListCmd := &cobra.Command{
		Use:   "list",
		Short: "List SSL certificate orders",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetSSLOrderList(&apisssl.GetSSLOrderListRequest{
				Options: common.StandQueryParameters{Page: 1, PerPage: 50},
			})
			if err != nil {
				return err
			}
			items := make([]model.SslOrder, len(resp.Data.Records))
			for i, o := range resp.Data.Records {
				items[i] = toSslOrder(o)
			}
			return (*out).Print(items)
		},
	}

	orderCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "Create an SSL certificate order",
		RunE: func(cmd *cobra.Command, args []string) error {
			req, err := orderFromFlags(cmd)
			if err != nil {
				return err
			}
			if _, err := (*rySDK).CreateSSLOrder(req); err != nil {
				return err
			}
			fmt.Println("SSL certificate order creating")
			return nil
		},
	}

	orderPriceCmd := &cobra.Command{
		Use:   "price",
		Short: "Get SSL certificate order price",
		RunE: func(cmd *cobra.Command, args []string) error {
			req, err := orderFromFlags(cmd)
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetSSLOrderPrice(req)
			if err != nil {
				return err
			}
			return (*out).Print(model.SslOrderPrice{
				Price: resp.Data.Price, Reward: resp.Data.Reward, RewardPoints: resp.Data.RewardPoints,
			})
		},
	}

	orderGetCmd := &cobra.Command{
		Use:   "get <id>",
		Short: "Get SSL certificate order details",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetSSLOrderDetail(id)
			if err != nil {
				return err
			}
			return (*out).Print(toSslOrderDetail(resp.Data))
		},
	}

	orderAssignCmd := &cobra.Command{
		Use:   "assign <id>",
		Short: "Add an order certificate to the certificate list",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).AssignSSLOrder(id); err != nil {
				return err
			}
			fmt.Printf("Order %s certificate assigned\n", args[0])
			return nil
		},
	}

	orderCertCmd := &cobra.Command{
		Use:   "cert <id>",
		Short: "Get the certificate of an order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			resp, err := (*rySDK).GetSSLOrderCert(id)
			if err != nil {
				return err
			}
			return (*out).Print(toSslCertDetail(resp.Data))
		},
	}

	orderDescriptionCmd := &cobra.Command{
		Use:   "description <id> <description>",
		Short: "Update an order description",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			if _, err := (*rySDK).UpdateSSLOrderDescription(id, args[1]); err != nil {
				return err
			}
			fmt.Printf("Order %s description updated\n", args[0])
			return nil
		},
	}

	orderRevokeCmd := &cobra.Command{
		Use:   "revoke <id>",
		Short: "Apply to revoke an SSL certificate",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			reason, _ := cmd.Flags().GetString("reason")
			letter, _ := cmd.Flags().GetString("letter")
			if reason == "" {
				return fmt.Errorf("--reason flag is required")
			}
			if _, err := (*rySDK).RevokeSSLOrder(id, reason, letter); err != nil {
				return err
			}
			fmt.Printf("Order %s revoking\n", args[0])
			return nil
		},
	}

	orderVerifyCmd := &cobra.Command{
		Use:   "verify <id>",
		Short: "Verify an SSL certificate order",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := cliutil.ParseID(args[0])
			if err != nil {
				return err
			}
			forceRefresh, _ := cmd.Flags().GetBool("force-refresh")
			if _, err := (*rySDK).VerifySSLOrder(id, forceRefresh); err != nil {
				return err
			}
			fmt.Printf("Order %s verified\n", args[0])
			return nil
		},
	}

	productsCmd := &cobra.Command{
		Use:   "products",
		Short: "List SSL certificate products",
		RunE: func(cmd *cobra.Command, args []string) error {
			resp, err := (*rySDK).GetSSLProductList()
			if err != nil {
				return err
			}
			items := make([]model.SslProduct, len(resp.Data))
			for i, p := range resp.Data {
				items[i] = toSslProduct(p)
			}
			return (*out).Print(items)
		},
	}

	applyCreateCmd.Flags().String("domains", "", "Domain(s) (required)")
	applyCreateCmd.Flags().String("verify-method", "", "Verification method: dns/http/auto (required)")
	addOrderFlags(orderCreateCmd)
	addOrderFlags(orderPriceCmd)
	orderRevokeCmd.Flags().String("reason", "", "Revoke reason (required)")
	orderRevokeCmd.Flags().String("letter", "", "Revoke letter content (Base64, required for non-DV)")
	orderVerifyCmd.Flags().Bool("force-refresh", false, "Force refresh the certificate")

	applyCmd.AddCommand(applyListCmd, applyCreateCmd, applyVerifyCmd)
	orderCmd.AddCommand(orderListCmd, orderCreateCmd, orderPriceCmd, orderGetCmd, orderAssignCmd, orderCertCmd, orderDescriptionCmd, orderRevokeCmd, orderVerifyCmd)
	sslCmd.AddCommand(listCmd, getCmd, uploadCmd, replaceCmd, deleteCmd, applyCmd, orderCmd, productsCmd)

	return sslCmd
}

func toSslCertRecord(r apisssl.SslCertificateRecord) model.SslCertRecord {
	return model.SslCertRecord{
		ID:     r.ID,
		Domain: r.Domain,
		Issuer: r.Issuer,
		Start:  formatDate(r.StartDate),
		Expire: formatDate(r.ExpDate),
	}
}

func toSslCertDetail(d apisssl.SslDetailData) model.SslCertDetail {
	return model.SslCertDetail{
		Domain: d.DomainName,
		Issuer: d.Issuer,
		Start:  formatDate(d.StartDate),
		Expire: formatDate(d.ExpDate),
		Remain: d.RemainDays,
	}
}

func formatDate(unix int) string {
	return time.Unix(int64(unix), 0).Format("2006-01-02")
}

func formatDate64(unix int64) string {
	return time.Unix(unix, 0).Format("2006-01-02")
}

func toSslOrder(o apisssl.SslOrder) model.SslOrder {
	return model.SslOrder{
		ID:     o.ID,
		Domain: orderDomains(o),
		Status: o.Status,
		Price:  o.Price,
		Expire: formatDate64(o.CertExpireAt),
	}
}

func toSslOrderDetail(o apisssl.SslOrder) model.SslOrderDetail {
	return model.SslOrderDetail{
		ID:      o.ID,
		Domain:  orderDomains(o),
		Status:  o.Status,
		Product: o.Product.Name,
		Expire:  formatDate64(o.CertExpireAt),
		Remain:  remainDays(o.CertExpireAt),
	}
}

func orderDomains(o apisssl.SslOrder) string {
	domains := append([]string{o.CsrInfo.CommonName}, o.CsrInfo.DNSNames...)
	return strings.Join(domains, ", ")
}

func remainDays(expire int64) string {
	d := int(time.Until(time.Unix(expire, 0)).Hours() / 24)
	return strconv.Itoa(d)
}

func toSslProduct(p apisssl.SslProduct) model.SslProduct {
	return model.SslProduct{
		ID:       p.ID,
		Name:     p.Name,
		Type:     p.Type,
		Brand:    p.Brand,
		Price12:  p.PriceMap["12"],
		Original: p.OriginalPriceMap["12"],
	}
}

func addOrderFlags(cmd *cobra.Command) {
	cmd.Flags().String("domains", "", "Domain(s) (required)")
	cmd.Flags().Int("duration", 0, "Duration in months (required)")
	cmd.Flags().Int("product-id", 0, "SSL product ID (required, see 'ry ssl products')")
	cmd.Flags().Float64("price", 0, "Price for verification")
	cmd.Flags().Int("coupon", 0, "Coupon ID")
}

func orderFromFlags(cmd *cobra.Command) (*apisssl.CreateSSLOrderRequest, error) {
	domains, _ := cmd.Flags().GetString("domains")
	duration, _ := cmd.Flags().GetInt("duration")
	productID, _ := cmd.Flags().GetInt("product-id")
	if domains == "" || duration == 0 || productID == 0 {
		return nil, fmt.Errorf("--domains, --duration and --product-id flags are required")
	}
	price, _ := cmd.Flags().GetFloat64("price")
	coupon, _ := cmd.Flags().GetInt("coupon")
	return &apisssl.CreateSSLOrderRequest{
		Domains: domains, Duration: duration, Price: price,
		ProductID: productID, WithCouponID: coupon,
	}, nil
}
