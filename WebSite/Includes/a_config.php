<?php
	switch ($_SERVER["SCRIPT_NAME"]) {
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
