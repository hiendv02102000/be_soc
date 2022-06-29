package utils

import (
	"context"
	"fmt"
	"io"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

var PUBLIC_FOLDER = "docs/sdk/go/"

func UploadImageToCloud(file io.Reader, id_file string, width int, height int) (url string, err error) {
	cld, err := cloudinary.NewFromParams("dtctadira", "658231379476276", "c8yMID5XRYNirCK5jjCCzuj0HE0")
	if err != nil {
		return
	}
	ctx := context.Background()
	publicId := PUBLIC_FOLDER + id_file
	// 3. Upload an image
	//===================

	_, err = cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: publicId,
		Transformation: "c_crop,g_center/q_auto/f_auto", Tags: []string{"img"}})
	if err != nil {
		return
	}
	// 4. Get your image information
	//===================
	my_image, err := cld.Image(publicId)
	if err != nil {
		fmt.Println("error")
	}

	// 5. Transform your image
	//=========================

	// Resize to 250 x 250 pixels using the 'fill' crop mode.
	my_image.Transformation = fmt.Sprintf("c_fill,h_%d,w_%d", height, width) //"c_fill,h_250,w_250"

	// 6. Generate your image URL
	//=========================

	url, err = my_image.String()
	return
}

func UploadFile(file io.Reader, id_file string) (url string, err error) {
	//cld, err := cloudinary.NewFromParams("hfbryanqg", "813372748842467", "IReNAP12GJOEbPay753ynjrQnS4")
	cld, err := cloudinary.NewFromParams("dtctadira", "658231379476276", "c8yMID5XRYNirCK5jjCCzuj0HE0")
	if err != nil {
		return
	}
	ctx := context.Background()
	publicId := PUBLIC_FOLDER + id_file
	res, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: publicId,
		Transformation: "c_crop,g_center/q_auto/f_auto", Tags: []string{"docx"}})
	url = res.URL
	return
}

// func GetURLFile(id_file string) (url string, err error) {
// 	cld, err := cloudinary.NewFromParams("hfbryanqg", "813372748842467", "IReNAP12GJOEbPay753ynjrQnS4")
// 	if err != nil {
// 		return
// 	}
// 	publicId := PUBLIC_FOLDER + id_file
// 	my_docx_file, err := cld.File(publicId)
// 	if err != nil {
// 		return
// 	}
// 	url, err = my_docx_file.String()
// 	return
// }
// func GetURLImg(id_file string, width int, height int) (url string, err error) {
// 	cld, err := cloudinary.NewFromParams("hfbryanqg", "813372748842467", "IReNAP12GJOEbPay753ynjrQnS4")
// 	if err != nil {
// 		return
// 	}
// 	publicId := PUBLIC_FOLDER + id_file
// 	my_image, err := cld.Image(publicId)
// 	if err != nil {
// 		fmt.Println("error")
// 	}

// 	// 5. Transform your image
// 	//=========================

// 	// Resize to 250 x 250 pixels using the 'fill' crop mode.
// 	my_image.Transformation = fmt.Sprintf("c_fill,h_%d,w_%d", height, width) //"c_fill,h_250,w_250"

// 	// 6. Generate your image URL
// 	//=========================

// 	url, err = my_image.String()
// 	return
// }
