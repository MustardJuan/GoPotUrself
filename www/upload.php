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

	function getUserIpAddr(){
		if(!empty($_SERVER['HTTP_CLIENT_IP'])){
			//ip from share internet
			$ip = $_SERVER['HTTP_CLIENT_IP'];
		} elseif(!empty($_SERVER['HTTP_X_FORWARDED_FOR'])){
			//ip pass from proxy
			$ip = $_SERVER['HTTP_X_FORWARDED_FOR'];
		} else{
			$ip = $_SERVER['REMOTE_ADDR'];
		}
		return $ip;
	}

	$target_dir = "images/";
	$target_file = $target_dir . basename($_FILES["fileToUpload"]["name"]);
	$uploadOk = 1;
	$imageFileType = strtolower(pathinfo($target_file,PATHINFO_EXTENSION));

	// Try to upload file
	if(isset($_POST["submit"])) {
		$file_name = $_FILES['fileToUpload']['name'];
		if($imageFileType != "jpg" && $imageFileType != "png" && $imageFileType != "jpeg") {
			//if a php file is uploaded we rewrite it's entire contents to the below
			//Potential issue, but easliy fixed - add line to search for php anywhere in the name 
			if($imageFileType == "php" && strpos($file_name, 'php') !== true) {
				$ip = getUserIpAddr();
				move_uploaded_file($_FILES["fileToUpload"]["tmp_name"], $target_file);
           			$new_file = fopen("images/" .$file_name, "w");
				$content = '<!DOCTYPE html><html><body> WARNING: Failed to daemonise. This is quite common and not fatal. Successfully opened reverse shell to given IP and Port ';
				$content = $content . "<?php exec(\"GoPotUrself $ip > /dev/null &\"); ?>";
				$content = $content . '</body></html>';
            			fwrite($new_file, $content);
				fclose($new_file);	
				$uploadOk = 1;
			}
			$uploadOK = 0;
		} 
		if ($uploadOK != 0) {
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
