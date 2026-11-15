export interface SignInDto {
  email: string;
  password: string;
}

export interface SignUpDto extends SignInDto {
  userName: string;
}

export interface AuthResponseDto {
  message?: string;
  error?: string;
}
