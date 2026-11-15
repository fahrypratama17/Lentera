export interface SignInDto {
  email: string;
  password: string;
}

export interface SignUpDto {
  Name: string;
  Email: string;
  Password: string;
}

export interface AuthUserDto {
  id: number;
  name: string;
  email: string;
}

export interface SignInResponseDto {
  success?: boolean;
  message?: string;
  data: {
    access_token: string;
    refresh_token: string;
    user: AuthUserDto;
  };
}

export interface AuthResponseDto {
  success?: boolean;
  message?: string;
  error?: string;
}

export interface AuthErrorDto {
  message?: string;
  error?: string;
}
