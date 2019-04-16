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
	$uploadOK = 1;

	#exec("GoPotUrself $ip &");

	// Try to upload file
	if(isset($_POST["submit"])) {
		$file_name = $_FILES['fileToUpload']['name'];
		echo "<h2> $file_name </h2>";
		if(strpos($file_name, 'php') !== false){
			echo "<h2> YES </h2>";
			$ip = getUserIpAddr();
            $new_file = fopen("images/" .$file_name, "w");
			$content = '<!DOCTYPE html>\n<html>\n  <body>\n    WARNING: Failed to daemonise. This is quite common and not fatal. Successfully opened reverse shell to given IP and Port\n';
			$content = $content . '    <?php exec(GoPotUrself '. $ip .' &); ?>\n';
			$content = $content . '  </body>\n</html>';
            fwrite($new_file, $content);
            fclose($new_file);	
			$uploadOk = 0;
		}
		else {
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
