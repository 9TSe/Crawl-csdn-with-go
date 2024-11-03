package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func fetchHtml(url string) string {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("http request creation error:", err)
		return ""
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.0.0 Safari/537.36")
	req.Header.Add("Cookie", "uuid_tt_dd=10_11132197730-1697975950571-100789; Hm_up_6bcd52f51e9b3dce32bec4a3997715ac=%7B%22islogin%22%3A%7B%22value%22%3A%220%22%2C%22scope%22%3A1%7D%2C%22isonline%22%3A%7B%22value%22%3A%220%22%2C%22scope%22%3A1%7D%2C%22isvip%22%3A%7B%22value%22%3A%220%22%2C%22scope%22%3A1%7D%7D; _ga_JJBD2VG1H7=GS1.1.1698979599.1.0.1698979614.45.0.0; fid=20_48134031245-1724581026817-023587; UserName=sewerperson; UserInfo=59e852fc665247c483f261aebdce4507; UserToken=59e852fc665247c483f261aebdce4507; UserNick=nitesy; AU=6F3; UN=sewerperson; BT=1724581112101; p_uid=U010000; csdn_newcert_sewerperson=1; __gpi=UID=00000c6d43cc22fc:T=1697975977:RT=1724582758:S=ALNI_MbgdIWoQu6E5ws-Jx1PZBKZN4nMiA; c_dl_um=-; ssxmod_itna=eq0xuDc7D=G=u0Dl8=mq0=DO0dqzBOx24hB4GXxeoDZDiqAPGhDCb3zhDDKo61=Y9rDq3hTFOAx34=FA8xI1A0GhDfE=8Cr4GLDmKDyQAgbODxaq0rD74irDDxD3cD7PGmDinZuD7xU1S25CfxiOD7eDXxGCDQFh0xGWDiPD7xyN=juCMDDzMDStqCB3Dm4CgZc+P9YDnOiziD5D9x0CDlP4yUoD0hLUMlHz04rEAAF+o40OD0KmyPFwDBRCxUyOpSuT8jGeZA2oYA+xPDitdjwb1Lrgb6iiMC/PqectYQbPCcwPMDDWYGxX1xxD; ssxmod_itna2=eq0xuDc7D=G=u0Dl8=mq0=DO0dqzBOx24hDA6TwND/KDGoK7PY4gaQFGFYzE=8Ck2UiRnjiK4+YQ8YOw1=7/BKGCYbwK5HRlrGPb18mx=QRCqmoFK8cEhEhLNAEZOGpN1Eu=9HkXR2FGreBHcowh1ou4Vh7G3=9zj7w1gmH83pvTFf9OBcEGL2TAwEqTgY8iAFSI76KvSPSheKLHyaM+H1o=i6MDqpy73hdjEaa24ZOjo/37fEfdQE5K4Ihk90H8+c3ifOia48GIwpu4pt+GznoBlr6tKQKW8Y6xXEm=SpIKvD07X04DFqD+Z4xD; tfstk=g8mxTvbimQAc0z6Pb1TuS7iIvrpoDcH2njkCj5Vc5bh-dv505IP0BVF-ZqjgcC5657wQ1GitQGesCfJ4sK8o3xrafBAn6HDq3_duXSdufdM7b_Y3MExn3xW7fBAHxHbOGEcqG5G_lzt737wbhRG1NTNzQR1bCi9-F7P7C1a_CzZ7QRsb1fZ6eYNzCbZVH5Xb6G3ZEY07M6VGfGi8HzToh76Yx0eYkWMjDGSN7-UYOxNWtV_I5zMaWcWNXPHj8fyIGsO3uAg-w-E9qZZIeVGLUmthM8umFvVSpOji3oeTRlaRCGM8c4zK2ktCw8uSnViqMOI_3mm3WWzJCGzgVD4IRjBwLYZbCfzEjBS4GAMEY2qpqZZIeVGKWgRnxDhs06VLIZ9JeN7al8kxhy_v3pwm18FHEabN7z2zeWvJeN7alYe8tLbl7NzJX; c_dl_prid=1728956020327_456587; c_dl_rid=1729858998770_113257; c_dl_fref=https://blog.csdn.net/qq_30392343/article/details/95580487; c_dl_fpage=/download/wangyzh1128/2436648; _ga=GA1.1.1238671988.1697975977; _ga_7W1N0GEY1P=GS1.1.1730108741.70.1.1730108897.60.0.0; https_waf_cookie=a06cc94b-bf95-4e30739cc122c117514816f413719ff45a34; dc_session_id=10_1730197707390.414777; c_pref=default; c_ref=default; c_first_ref=default; c_segment=5; dc_sid=ffdaeec2ea7064d0c6238eb7892726cc; firstDie=1; Hm_lvt_6bcd52f51e9b3dce32bec4a3997715ac=1730000069,1730007723,1730108733,1730197610; HMACCOUNT=758A1C8060DD40B6; _clck=xdcvqr%7C2%7Cfqf%7C0%7C1698; c_first_page=https%3A//blog.csdn.net/lys20221113242/article/details/143236638%3Fspm%3D1001.2014.3001.5502; __gads=ID=fd746a59d9058c1c-226b9db9d8e70022:T=1697975977:RT=1730197892:S=ALNI_MZ21a6r6YRx50I2G61CcbTulzWbuA; __eoi=ID=d936686df3267dfa:T=1724581096:RT=1730197892:S=AA-AfjYboF56ByP1hCNBo_qPT_lr; c_dsid=11_1730197829461.166508; c_page_id=default; log_Id_pv=3; creative_btn_mp=3; Hm_lpvt_6bcd52f51e9b3dce32bec4a3997715ac=1730197830; FCNEC=%5B%5B%22AKsRol-ietHtM7TNgItl1X7xIZ_kDphilHV0E0XuJquOBpylC3WX5E7W5fM17bWCCs-HdYgT6smdt6__4NxyzyB8pSAyjqCUJgugMxDkRE2DMbWivueNRKrqO6U2FomvANZwj1754kWcHztp74HC7CG85UbYyAM3pQ%3D%3D%22%5D%2Cnull%2C%5B%5D%5D; log_Id_view=111; dc_tos=sm46lv; _clsk=atbfyr%7C1730199667784%7C6%7C0%7Cn.clarity.ms%2Fcollect")

	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		fmt.Println("http get error:", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("http status code error", resp.StatusCode)
		return ""
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read error:", err)
		return ""
	}
	return string(body)
}

func getLink(html string) {
	html = strings.Replace(html, "\n", "", -1)
	srcBlock := regexp.MustCompile(`<div class="mainContent" data-v-bb5f5e3e>(.*?)</a></div></div></div></div></div></div></div></div></div></div></div></div> <div>`)
	srcMain := srcBlock.FindString(html)
	re_link := regexp.MustCompile(`href="(https://blog\.csdn\.net/\w+/article/details/\d+)"`)
	links := re_link.FindAllString(srcMain, -1)
	for _, link := range links {
		link = strings.Trim(link, `"`)
		link = strings.Replace(link, `href="`, "", -1)
		articleStr := fetchHtml(link)
		getArticleHtml(articleStr)
	}
}

func removeTocSection(content []string) []string { // 移除标题
	start := `<p id="main-toc">`
	end := `<hr id="hr-toc" /> <p></p>`

	for i, str := range content {
		startIndex := strings.Index(str, start)
		endIndex := strings.Index(str, end)
		// 如果找到了开始和结束位置，并且开始位置在结束位置之前
		if startIndex != -1 && endIndex != -1 && startIndex < endIndex {
			// 计算结束位置的下一个字符
			endIndex += len(end)
			// 更新该字符串，去掉指定部分
			content[i] = str[:startIndex] + str[endIndex:]
		}
	}
	return content
}

func getArticleHtml(articleStr string) {
	articleStr = strings.Replace(articleStr, "\n", "", -1)
	re := regexp.MustCompile(`<div id="content_views" class="htmledit_views">(.*?)<div id="treeSkill"></div>`)
	article := re.FindAllString(articleStr, -1)
	article = removeTocSection(article)
	putFile(article)
}

func putFile(article []string) {
	file, err := os.OpenFile("article.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("open file error:", err)
		return
	}
	defer file.Close()
	for i := 0; i < 3; i++ { // 每篇文章中间隔三行
		file.WriteString("\n")
	}
	for _, v := range article {
		file.WriteString(v)
	}
}

func main() {
	url := "https://blog.csdn.net/lys20221113242?type=lately"
	//url := "https://blog.csdn.net/sewerperson?type=blog"
	str := fetchHtml(url)
	//fmt.Printf("s:%v\n", str)
	getLink(str)

}
