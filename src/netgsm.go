package netgsm

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Config struct {
	ApiUrl       string
	SmsCompany   string
	SmsMsgHeader string
	SmsUserCode  string
	SmsPassword  string
}

type API struct {
	Config Config
}

type SendSmsRequest struct {
	XMLName  xml.Name `xml:"xml,omitempty"`
	MainBody struct {
		Header struct {
			Company  string `xml:"company,omitempty"`
			UserCode string `xml:"usercode,omitempty"`
			Password string `xml:"password,omitempty"`
			// Gönderime başlayacağınız tarih. (ddMMyyyyHHmm)* Boş bırakılırsa mesajınız hemen gider.
			StartDate string `xml:"startdate,omitempty"`
			// İki tarih arası gönderimlerinizde bitiş tarihi.(ddMMyyyyHHmm) * Boş bırakılırsa sistem başlangıç tarihine 21 saat ekleyerek otomatik gönderir..
			StopDate string `xml:"stopdate,omitempty"`
			// 1:N gönderimlerde <type>1:n</type> olarak n:N gönderimlerde <type>n:n</type> olarak teslim edebilirsiniz.
			Type      string `xml:"type,omitempty"`
			MsgHeader string `xml:"msgheader,omitempty"`
		} `xml:"header,omitempty"`
		Body struct {
			Msg string `xml:"msg,omitempty"`
			No  string `xml:"no,omitempty"`
		} `xml:"body,omitempty"`
	} `xml:"mainbody,omitempty"`
}

func (api *API) Sms(request SendSmsRequest) (bool, error) {
	errorCodes := make(map[string]string)
	errorCodes["30"] = "Geçersiz kullanıcı adı , şifre veya kullanıcınızın API erişim izninin olmadığını gösterir. Ayrıca eğer API erişiminizde IP sınırlaması yaptıysanız ve sınırladığınız ip dışında gönderim sağlıyorsanız 30 hata kodunu alırsınız. API erişim izninizi veya IP sınırlamanızı , web arayüzden; sağ üst köşede bulunan ayarlar> API işlemleri menüsunden kontrol edebilirsiniz."
	errorCodes["20"] = "Mesaj metninde ki problemden dolayı gönderilemediğini veya standart maksimum mesaj karakter sayısını geçtiğini ifade eder.(Standart maksimum karakter sayısı 917 dir. Eğer mesajınız türkçe karakter içeriyorsa Türkçe Karakter Hesaplama menüsunden karakter sayılarının hesaplanış şeklini görebilirsiniz.)"
	errorCodes["40"] = "Mesaj başlığınızın (gönderici adınızın) sistemde tanımlı olmadığını ifade eder. Gönderici adlarınızı API ile sorgulayarak kontrol edebilirsiniz."
	errorCodes["50"] = "Abone hesabınız ile İYS kontrollü gönderimler yapılamamaktadır."
	errorCodes["51"] = "Aboneliğinize tanımlı İYS Marka bilgisi bulunamadığını ifade eder"
	errorCodes["70"] = "Hatalı sorgulama. Gönderdiğiniz parametrelerden birisi hatalı veya zorunlu alanlardan birinin eksik olduğunu ifade eder"
	errorCodes["80"] = "Gönderim sınır aşımı"
	errorCodes["85"] = "Mükerrer Gönderim sınır aşımı. Aynı numaraya 1 dakika içerisinde 20'den fazla görev oluşturulamaz"
	errorCodes["100"] = "Sistem hatası"
	errorCodes["101"] = "Sistem hatası"

	request.MainBody.Header.Company = api.Config.SmsCompany
	request.MainBody.Header.MsgHeader = api.Config.SmsMsgHeader
	request.MainBody.Header.UserCode = api.Config.SmsUserCode
	request.MainBody.Header.Password = api.Config.SmsPassword
	request.MainBody.Header.Type = "1:n"
	request.MainBody.Body.Msg = "<![CDATA[" + request.MainBody.Body.Msg + " - ]]>"
	postdata, err := xml.Marshal(request)
	if err != nil {
		return false, err
	}
	rpl := strings.NewReplacer("&lt;!", "<!", "]&gt;", "]>", "<xml>", "", "</xml>", "")
	resp, err := http.Post(api.Config.ApiUrl, "text/xml; charset=utf-8", strings.NewReader(xml.Header+rpl.Replace(string(postdata))))
	if err != nil {
		return false, err
	}

	fmt.Println("Status Code :", resp.StatusCode)

	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		if errorCodes[bodyString] != "" {
			return false, errors.New(errorCodes[bodyString])
		}
	} else {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		return false, errors.New(bodyString)
	}

	defer resp.Body.Close()
	return true, nil
}

type IncommingSmsRequest struct {
	XMLName  xml.Name `xml:"xml,omitempty"`
	MainBody struct {
		Header struct {
			UserCode string `xml:"usercode,omitempty"`
			Password string `xml:"password,omitempty"`
			// Gönderime başlayacağınız tarih. (ddMMyyyyHHmm)* Boş bırakılırsa mesajınız hemen gider.
			StartDate string `xml:"startdate,omitempty"`
			// İki tarih arası gönderimlerinizde bitiş tarihi.(ddMMyyyyHHmm) * Boş bırakılırsa sistem başlangıç tarihine 21 saat ekleyerek otomatik gönderir..
			StopDate string `xml:"stopdate,omitempty"`
			// 1:N gönderimlerde <type>1:n</type> olarak n:N gönderimlerde <type>n:n</type> olarak teslim edebilirsiniz.
			Type string `xml:"type,omitempty"`
		} `xml:"header,omitempty"`
	} `xml:"mainbody,omitempty"`
}

type IncomingMessage struct {
	PhoneNumber string    `json:"phone_number"`
	Message     string    `json:"message"`
	MessageTime time.Time `json:"message_time"`
}

func (api *API) IncommingSms(request IncommingSmsRequest) ([]IncomingMessage, error) {
	errorCodes := make(map[string]string)
	errorCodes["30"] = "Geçersiz kullanıcı adı , şifre veya kullanıcınızın API erişim izninin olmadığını gösterir. Ayrıca eğer API erişiminizde IP sınırlaması yaptıysanız ve sınırladığınız ip dışında gönderim sağlıyorsanız 30 hata kodunu alırsınız. API erişim izninizi veya IP sınırlamanızı , web arayüzden; sağ üst köşede bulunan ayarlar> API işlemleri menüsunden kontrol edebilirsiniz."
	errorCodes["40"] = "Gösterilecek mesajınızın olmadığını ifade eder. Api ile mesajlarınızı eğer startdate ve stopdate parametlerini kullanmıyorsanız sadece bir kere listeyebilirsiniz. Listelenen mesajlar diğer sorgulamalarınızda gelmez"
	errorCodes["50"] = "Tarih formatı hatalıdır. (Tarih formatı: ddmmyyyyhhmm şeklinde olmalıdır.)"
	errorCodes["60"] = "Arama kiterlerindeki startdate ve stopdate zaman farkının 30 günden fazla olduğunu ifade eder"
	errorCodes["70"] = "Hatalı sorgulama. Gönderdiğiniz parametrelerden birisi hatalı veya zorunlu alanlardan birinin eksik olduğunu ifade eder"

	request.MainBody.Header.UserCode = api.Config.SmsUserCode
	request.MainBody.Header.Password = api.Config.SmsPassword
	request.MainBody.Header.Type = "10000"
	postdata, err := xml.Marshal(request)
	if err != nil {
		return nil, err
	}
	rpl := strings.NewReplacer("&lt;!", "<!", "]&gt;", "]>", "<xml>", "", "</xml>", "")
	resp, err := http.Post(api.Config.ApiUrl, "text/xml; charset=utf-8", strings.NewReader(xml.Header+rpl.Replace(string(postdata))))
	if err != nil {
		return nil, err
	}

	fmt.Println("Status Code :", resp.StatusCode)

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	defer resp.Body.Close()

	if errorCodes[bodyString] != "" {
		return nil, errors.New(errorCodes[bodyString])
	}

	var messages []IncomingMessage

	messagesStrArr := strings.Split(bodyString, "<br>")
	for _, m := range messagesStrArr {
		messageStrArr := strings.Split(m, " | ")
		if len(messageStrArr) > 1 {
			message := IncomingMessage{}
			message.PhoneNumber = messageStrArr[0]
			message.Message = messageStrArr[1]
			fmt.Println(messageStrArr[2])
			// 08.05.2021 21:55:37
			message.MessageTime, _ = time.Parse("02.01.2006 15:04:05", messageStrArr[2])
			messages = append(messages, message)
		}
	}

	return messages, nil
}
