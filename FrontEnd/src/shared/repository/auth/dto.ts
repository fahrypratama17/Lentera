export interface SignInDto {
  email: string;
  password: string;
}

export interface SignUpDto {
  Name: string;
  Email: string;
  Password: string;
}

export interface AuthResponseDto {
  message?: string;
  error?: string;
}

export interface AuthErrorDto {
  message?: string;
  error?: string;
}
