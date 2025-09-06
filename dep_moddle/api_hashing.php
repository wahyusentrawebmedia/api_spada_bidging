<?php
header('Content-Type: application/json');

// Ambil input JSON dari request body
$input = json_decode(file_get_contents('php://input'), true);

if (!isset($input['password'])) {
    http_response_code(400);
    echo json_encode(['error' => 'Password is required']);
    exit;
}

$password = $input['password'];
$hash = password_hash($password, PASSWORD_DEFAULT);

echo json_encode(['hash' => $hash]);