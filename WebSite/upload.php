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
	
	// Try to upload file
	if(isset($_POST["submit"])) {
		
		if($imageFileType != "jpg" && $imageFileType != "png" && $imageFileType != "jpeg" && $imageFileType != "php" ) {
    			$uploadOk = 0;
		}	
   		if (move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $target_file) && $uploadOk != 0) {
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
