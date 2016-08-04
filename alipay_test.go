package alipay

import (
	"fmt"
	"github.com/Centny/gwf/routing/httptest"
	"testing"
)

func TestAlipay(t *testing.T) {
	var client = NewClient("https://mapi.alipay.com/gateway.do", nil)
	fmt.Println(client.CreateWebUrl("http://pb.dev.jxzy.com/_echo_", "http://pb.dev.jxzy.com/_echo_", "6843192280647119", "abcc", "223", 0.01))
	var ts = httptest.NewMuxServer()
	ts.Mux.HFunc("^/return(\\?.*)?$", client.Return)
	fmt.Println(ts.G("/return?%v", "body=223&buyer_email=centny%40gmail.com&buyer_id=2088102972036594&exterface=create_direct_pay_by_user&is_success=T&notify_id=RqPnCoPT3K9%252Fvwbh3InWfjSquPZ53GKZDlpLiPerRyczkZ1BqSCeryalHBnmC%252FQ3uhhI&notify_time=2016-08-04+11%3A15%3A02&notify_type=trade_status_sync&out_trade_no=6843192280647112&payment_type=1&seller_email=itdayang%40gmail.com&seller_id=2088501949844011&subject=abcc&total_fee=0.01&trade_no=2016080421001004590289703858&trade_status=TRADE_SUCCESS&sign=f98956240273d3bda99b84c9a64c27a4&sign_type=MD5"))
}
