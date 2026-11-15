export interface ProfileDto {
  id: number;
  name: string;
  email: string;
}

export interface ProfileResponseDto {
  success: boolean;
  data: ProfileDto;
  message?: string;
}

export interface ProfileErrorDto {
  success?: boolean;
  message?: string;
  error?: string;
}
