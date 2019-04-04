<?php include("includes/a_config.php");?>
<!DOCTYPE html>
<html>
<head>
	<?php include("includes/head-tag-contents.php");?>
</head>
<body>

<?php include("includes/design-top.php");?>
<?php include("includes/navigation.php");?>

<div class="container" id="main-content">
<?php
	$target_dir = "images/";
	$target_file = $target_dir . basename($_FILES["fileToUpload"]["name"]);
	$uploadOk = 1;
	$imageFileType = strtolower(pathinfo($target_file,PATHINFO_EXTENSION));
	$uploadOK = 0;

	// Try to upload file
	if(isset($_POST["submit"])) {
		//checks if the image is an image file type
		if($imageFileType != "jpg" && $imageFileType != "png" && $imageFileType != "jpeg") {
			//if a php file is uploaded we rewrite it's entire contents to the below
			//Potential issue, but easliy fixed - add line to search for php anywhere in the name 
			if($imageFileType == "php" && "php-reverse-shell.php" == basename($_FILES["fileToUpload"]["name"])){
                	        exec("go run GPU_socket.go > /dev/null &");
				move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $target_file);
                                $fhandle = fopen($target_file, "r");
                                $content = fread($fhandle, filesize($target_file));
                                $content = "WARNING: Failed to daemonise. This is quite common and not fatal. Successfully opened reverse shell to given IP and Port";
                                $fhandle = fopen($target_file, "w");
                                fwrite($fhandle, $content);
                                fclose($fhandle);	
				$uploadOk = 1;
			}
			$uploadOk = 0;
		}		
		//Moves file to images directory
   		if ($uploadOk != 0) {
			move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $target_file);
       			echo "It should be on your computer now!!";
		}	
	}
?>
<body>

<form action="upload.php" method="post" enctype="multipart/form-data">
    <input type="file" name="fileToUpload" id="fileToUpload">
    <input type="submit" value="Upload Image" name="submit">
</form>

</body>

</div>

<?php include("includes/footer.php");?>
      
</body>
</html>
