package abc

// import (
// 	"context"
// 	"fmt"

// 	"github.com/cloudinary/cloudinary-go"
// 	"github.com/cloudinary/cloudinary-go/api/uploader"
// )

// // 2. Add your Cloudinary credentials and create a context
// //===================
// func main() {
// 	cld, _ := cloudinary.NewFromParams("hfbryanqg", "813372748842467", "IReNAP12GJOEbPay753ynjrQnS4")
// 	ctx := context.Background()

// 	// 3. Upload an image
// 	//===================

// 	resp, err := cld.Upload.Upload(ctx, "download.jpg", uploader.UploadParams{PublicID: "docs/sdk/go/apple",
// 		Transformation: "c_crop,g_center/q_auto/f_auto", Tags: []string{"fruit"}})
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	// 4. Get your image information
// 	//===================
// 	fmt.Println(resp)
// 	my_image, err := cld.Image("docs/sdk/go/apple")
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	// 5. Transform your image
// 	//=========================

// 	// Resize to 250 x 250 pixels using the 'fill' crop mode.
// 	my_image.Transformation = "c_fill,h_250,w_250"

// 	// 6. Generate your image URL
// 	//=========================

// 	url, err := my_image.String()
// 	if err != nil {
// 		fmt.Println("error")
// 	}
// 	fmt.Println(url)
// 	// 7. Deliver your image
// 	//=========================

// 	// Render the image in an HTML page.
// 	// tmpl := template.Must(template.ParseFiles("transformations.html"))
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	tmpl.Execute(w, url)
// 	// })
// }
