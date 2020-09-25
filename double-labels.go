type abType struct {
	Product           string
	Description       string
	UOM               string
	StorageBin        string
	StorageBinDisplay string
	Check1            string
	Check2            string
	TodayDate         string
	DescHeight        string
	DescWidth         string
	LongDesc1         string
	LongDesc2         string
	LongDescHeight    string
	LongDescWidth     string
	Arrow             string
}

type largeLabelType struct {
	A     abType
	B     abType
	ZPL   string
	Count int
}

func parseLabelLarge(t **template.Template) {
	zpl := `^XA
^FT265,1250
^A@B,70,65,E:ROB002.TTF^FD{{.A.Product}}^FS
^FT185,1550
^A@B,190,135,E:ROB002.TTF^FD{{.A.StorageBinDisplay}}^FS
^FT200,310
^A@B,230,,E:ROB002.TTF^FD{{.A.Check1}}^FS
^FT200,160
^A@B,230,,E:ROB002.TTF^FD{{.A.Check2}}^FS
^FT20,320
^GB195,150,2^FS
^FT20,172
^GB195,150,2^FS
^FO20,365
^BY3^BCB,260,N,,,A
^FD{{.A.StorageBin}}^FS
^FT375,1250
^A@B,{{.A.DescHeight}},{{.A.DescWidth}},E:ROB004.TTF^FB1245,1,0,C^FD{{.A.Description}}^FS
^FT115,340
^A@N,20,18,E:ROB002.TTF^FD{{.A.TodayDate}}^FS
^FT240,1425
^GB150,150,2^FS
^FT365,1440
^A@B,105,100,E:ROB002.TTF^FB180,1,0,C^FD{{.A.UOM}}^FS
^FT260,1555
^A@N,200,,E:SYM003.TTF^FD{{.A.Arrow}}^FS
{{ if eq .Count 2}}^FT680,1250
^A@B,70,65,E:ROB002.TTF^FD{{.B.Product}}^FS
^FT600,1550
^A@B,190,135,E:ROB002.TTF^FD{{.B.StorageBinDisplay}}^FS
^FT615,310
^A@B,230,,E:ROB002.TTF^FD{{.B.Check1}}^FS
^FT615,160
^A@B,230,,E:ROB002.TTF^FD{{.B.Check2}}^FS
^FT435,320
^GB195,150,2^FS
^FT435,172
^GB195,150,2^FS
^FO435,365
^BY3^BCB,260,N,,,A
^FD{{.B.StorageBin}}^FS
^FT790,1250
^A@B,{{.B.DescHeight}},{{.B.DescWidth}},E:ROB004.TTF^FB1245,1,0,C^FD{{.B.Description}}^FS
^FT530,340
^A@N,20,18,E:ROB002.TTF^FD{{.B.TodayDate}}^FS
^FT655,1425
^GB150,150,2^FS
^FT780,1440
^A@B,105,100,E:ROB002.TTF^FB180,1,0,C^FD{{.B.UOM}}^FS
^FT680,1555
^A@N,200,,E:SYM003.TTF^FD{{.B.Arrow}}^FS
{{end}}^XZ
`

	tt := template.New("t")
	tt, err := tt.Parse(zpl)
	if err != nil {
		fmt.Println(err)
	}
	//log.Println("check tt:", tt == nil)
	*t = tt

}
