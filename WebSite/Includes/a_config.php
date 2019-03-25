<?php
	switch ($_SERVER["SCRIPT_NAME"]) {
		case "/php-template/upload.php":
			$CURRENT_PAGE = "Uploads"; 
			$PAGE_TITLE = "Whoo buddy";
			break;		
		case "/php-template/pictures.php":
			$CURRENT_PAGE = "Pictures"; 
			$PAGE_TITLE = "Pics";
			break;
		case "/php-template/FirstPage.php":
			$CURRENT_PAGE = "FirstPage"; 
			$PAGE_TITLE = "First Page";
			break;
		case "/php-template/contact.php":
			$CURRENT_PAGE = "Contact"; 
			$PAGE_TITLE = "Contact Us";
			break;
		default:
			$CURRENT_PAGE = "Index";
			$PAGE_TITLE = "Welcome to my homepage!";
	}
?>
