<?php
require_once('../../config.php');
require_once($CFG->libdir . '/externallib.php');

// Enable CORS
header('Access-Control-Allow-Origin: *');
header('Access-Control-Allow-Methods: POST, OPTIONS');
header('Access-Control-Allow-Headers: Content-Type, Authorization, X-API-Key');

// Handle preflight
if ($_SERVER['REQUEST_METHOD'] == 'OPTIONS') {
    http_response_code(200);
    exit();
}

// Check API key (optional security layer)
$api_key = isset($_SERVER['HTTP_X_API_KEY']) ? $_SERVER['HTTP_X_API_KEY'] : '';
$valid_api_key = get_config('local_validatepassword', 'apikey');

if ($valid_api_key && $api_key !== $valid_api_key) {
    http_response_code(401);
    echo json_encode(['error' => 'Invalid API key']);
    exit();
}

// Only accept POST requests
if ($_SERVER['REQUEST_METHOD'] != 'POST') {
    http_response_code(405);
    echo json_encode(['error' => 'Method not allowed']);
    exit();
}

// Get and validate input
$input = json_decode(file_get_contents('php://input'), true);

if (!isset($input['userid']) || !isset($input['oldpassword'])) {
    http_response_code(400);
    echo json_encode(['error' => 'Missing required parameters']);
    exit();
}

$userid = clean_param($input['userid'], PARAM_INT);
$oldpassword = clean_param($input['oldpassword'], PARAM_RAW);

// Validate parameters
if ($userid <= 0 || empty($oldpassword)) {
    http_response_code(400);
    echo json_encode(['error' => 'Invalid parameters']);
    exit();
}

try {
    global $DB;
    
    // Get user data
    $user = $DB->get_record('user', ['id' => $userid, 'deleted' => 0]);
    
    if (!$user) {
        http_response_code(404);
        echo json_encode([
            'valid' => false,
            'message' => 'User not found or deleted'
        ]);
        exit();
    }
    
    // Check if user is suspended
    if ($user->suspended) {
        http_response_code(403);
        echo json_encode([
            'valid' => false,
            'message' => 'User account suspended'
        ]);
        exit();
    }
    
    // Validate password
    if (validate_user_password($user, $oldpassword)) {
        $response = [
            'valid' => true,
            'message' => 'Password valid',
            'user' => [
                'id' => $user->id,
                'username' => $user->username,
                'email' => $user->email,
                'firstname' => $user->firstname,
                'lastname' => $user->lastname
            ]
        ];
    } else {
        $response = [
            'valid' => false,
            'message' => 'Password tidak valid'
        ];
    }
    
    http_response_code(200);
    echo json_encode($response);
    
} catch (Exception $e) {
    http_response_code(500);
    echo json_encode([
        'error' => 'Internal server error',
        'message' => $e->getMessage()
    ]);
}
?>