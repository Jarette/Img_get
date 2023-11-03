/*****************************************************************************
*                    
*  Author:           Jarette Greene
*  Email:            jkgreene0406@my.msutexas.edu / jarettegreene09@gmail.com
*  Label:            P03
*  Title:			 Image Ascii Art           
*  Course:           CMPS 4143
*  Semester:         Fall 2023
* 
*  Description:
*	
*		This Package contains functions that allow the user to enter a URL 
*		and download an image to thier system.
*
* 
*  Usage:
*    - import to main.go file 
* 
*  Files           
*       N/A 
*****************************************************************************/

package Img_get

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Getpic gets picture from the internet
func Getpic(url string) {
	// URL of the image you want to download
	//imageUrl := "https://images.pexels.com/photos/17929271/pexels-photo-17929271/free-photo-of-woman-standing-on-vineyard.jpeg"
	//imageURL := "https://s3e8p5g8.rocketcdn.me/wp-content/uploads/2020/11/midwestern-state-university2.jpg"
	imageURL:= url;
	// Create an HTTP GET request
	response, err := http.Get(imageURL)
	if err != nil {
		fmt.Println("Error making the request:", err)
		return
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		fmt.Println("Error: Status code", response.StatusCode)
		return
	}

	// Create a new file to save the image
	outputFile, err := os.Create("downloaded_image.jpg")
	if err != nil {
		fmt.Println("Error creating the file:", err)
		return
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		fmt.Println("Error saving the image:", err)
		return
	}
	// final message to say that the image was saved 
	fmt.Println("Image downloaded and saved as 'downloaded_image.jpg'")
}
