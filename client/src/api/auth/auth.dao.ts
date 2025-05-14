export interface RegisterRequest {
    email: string;
    password: string;
    full_name?: string;
    date_of_birth?: string;
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
    refresh_token: string;
    user_id: number;
    expire_at: number;
}
export interface RefreshTokenRequest {
    refresh_token?: string;
}
export interface RefreshTokenResponse {
    access_token: string;
    refresh_token: string;
    user_id: number;
    expire_at: number;
}
export interface LogoutRequest {
    refresh_token?: string;
}
export interface LogoutResponse {
    success: boolean;
}
