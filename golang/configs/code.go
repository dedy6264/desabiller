package configs

var (
	VALIDATE_ERROR_CODE = "34"
	DB_ERROR            = "81"
	DB_NOT_FOUND        = "82"

	//trx
	SUCCESS_MSG = "SUCCESS"
	PENDING_MSG = "PENDING"
	FAILED_MSG  = "FAILED"

	SUCCESS_CODE         = "00"
	PENDING_CODE         = "05"
	FAILED_CODE          = "09"
	INQUIRY_SUCCESS_CODE = "10"
	INQUIRY_PENDING_CODE = "15"
	INQUIRY_FAILED_CODE  = "19"
	//status code worker/helper
	WORKER_SUCCESS_CODE = "000"
	WORKER_PENDING_CODE = "500"

	WORKER_FAILED_CODE      = "900"
	WORKER_INVALID_PARAM    = "910"
	WORKER_CREDENTIAL_ERROR = "920"
	WORKER_VALIDATION_ERROR = "930"
	WORKER_SYSTEM_ERROR     = "940"
	// WORKER_VALIDATE_SYSTEM   = "090"
)

// failed{
// 	WORKER_FAILED_CODE
// 	06	TRANSACTION NOT FOUND	Failed	There is no transaction with your inputted ref_id. Please check again your inputted ref_id to find your transaction.
// 07	FAILED	Failed	Your current transaction has failed. Please try again.
// 13	CUSTOMER NUMBER BLOCKED	Failed	Your customer number (customer_id) has been blocked. You can change your customer number (customer_id) or contact our Customer Service.
// "18",	NUMBER NOT AVAILABLE	Failed	You can see all available E-SIM number by using E-SIM List API.
// "20",	CODE NOT FOUND	Failed	Your inputted product_code isn’t in the database. Check again your product_code, you can check product_code list by using Pricelist API.
// "21",	NUMBER EXPIRED	Failed	Your phone number (customer_id) is expired. You can try other phone number.
// "132",	PRODUCT CODE NOT ELIGIBLE DUE TO SUBSCRIBER LOCATION	Failed	Your inputted product_code isn’t eligible due to subscriber location. Please try again with different product_code.
// 	"106",	PRODUCT IS TEMPORARILY OUT OF SERVICE	Failed	The product_code that you pick is in non-active status. You can retry your transaction with another product_code that has active status.

// WORKER_VALIDATION_ERROR
// "14",	INCORRECT DESTINATION NUMBER	Failed	Your customer_id that you’ve inputted isn’t a valid phone number. Please check again your customer_id.
// "16",	NUMBER NOT MATCH WITH OPERATOR	Failed	Your phone number (customer_id) that you’ve inputted doesn’t match with your desired operator (product_code). Please check again your phone number or change your operator.
// "19",	NUMBER IS ALREADY IN USE	Failed	Please select other E-SIM number.
// "131",	TOP UP REGION BLOCKED FOR PLAYER	Failed	Your current destination number top up request is blocked in that region. Please try again with a different destination number.
// "141",	INVALID USER ID / ZONE ID / SERVER ID / ROLENAME	Failed	Your inputted user ID / Zone ID / Server ID / Role name isn’t valid. Please try again with another user ID / Zone ID / Server ID / Role name. You can check on Inquiry Game Server.
// "142",	INVALID USER ID	Failed	Your current destination number (user id) top up request is invalid. Please try again with a different destination number or try checking for typos in your field.
// "206",	THIS DESTINATION NUMBER HAS BEEN BLOCKED	Failed	The customer_id that you inputted is blocked or not in whitelist. You can unblock it by remove customer number blacklist in API Security menu blacklist (https://developer.iak.id/end-user-blacklist) or add customer number whitelist in API Security menu whitelist (https://developer.iak.id/end-user-whitelist) on developer.iak.id.

// WORKER_INVALID_PARAM
// "203",	NUMBER IS TOO LONG	Failed	Your inputted customer ID is too long. Please check again your customer ID.
// 	"205",	WRONG COMMAND	Failed	The command that you’ve inputted is not a valid command, try check your commands field for typos or try another command.
// 	"107",	ERROR IN XML FORMAT	Failed	The body format of your request isn’t correct or there is an error in your body (required, ajax error, etc). Please use the correct JSON or XML format corresponding to your request to API. You can see the required body request for each request in the API Documentation.

// WORKER_CREDENTIAL_ERROR
// 	"102",	INVALID IP ADDRESS	Failed	Your IP address isn’t allowed to make a transaction. You can add your IP address to your allowed IP address list in https://developer.iak.id/prod-setting.

// system
// 	"12",	BALANCE MAXIMUM LIMIT EXCEEDED	Failed	-
// 	"204",	WRONG AUTHENTICATION	Failed	Your sign (signature) field doesn’t contain the right key for your current request. Please check again your sign value.
// 	"17",	INSUFFICIENT DEPOSIT	Failed	Your current deposit is lower than the product_price you want to buy. You can add more money into your deposit by doing top up on iak.id deposit menu, or if you are in development mode, you can add your development deposit by clicking the + (plus) sign on development deposit menu (https://developer.iak.id/sandbox-report).
// 	"110",	SYSTEM UNDER MAINTENANCE	Failed	The system is currently under maintenance, you can try again later.
// 	"202",	MAXIMUM 1 NUMBER 1 TIME IN 1 DAY	Failed	You can only top up to a phone number once in a day (based on your developer setting). If you want to allow more than one top up to a phone number, you can go to https://developer.iak.id/ then choose API Setting menu, you can turn on “Allow multiple transactions for the same number” in development or production settings.
// 	"207",	MAXIMUM 1 NUMBER WITH ANY CODE 1 TIME IN 1 DAY	Failed	You’ve already done a transaction today. Please do another transaction tomorrow, or disable the high restriction setting in https://developer.iak.id/prod-setting.
// 	"121",	MONTHLY TOP UP LIMIT EXCEEDED	Failed	This response code applies to OVO products.
// 	"117",	PAGE NOT FOUND	Failed	The API URL that you want to hit is not found. Try checking your request URL for any typos or try other API URLs.
// 	"10",	REACH TOP UP LIMIT USING SAME DESTINATION NUMBER IN 1 DAY	Failed	Your current destination number top up request is reaching the limit on that day. Please try again tomorrow.

// }
