export interface RegisterRequest {
    email: string;
    password: string;
    full_name: string;
    date_of_birth: string;
}

export interface RegisterResponse {
    user_id: string;
}

export interface LoginRequest {
    email: string;
    password: string;
}
export interface LoginResponse {
    access_token: string;
    expire_at: number;
}
export interface RefreshTokenRequest {
    // No fields needed as session token comes from cookie
}
export interface RefreshTokenResponse {
    access_token: string;
    expire_at: number;
}
export interface LogoutRequest {
    // No fields needed as session token comes from cookie
}
export interface LogoutResponse {
    success: boolean;
    message: string;
}
export interface LogoutAllRequest {
    // No fields needed as user ID comes from JWT token
}
export interface LogoutAllResponse {
    success: boolean;
    message: string;
    sessions_ended: number;
}
