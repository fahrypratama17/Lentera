import { api } from "@/shared/lib/axios";

import type {
  SignInDto,
  SignUpDto,
  AuthResponseDto,
} from "@/shared/repository/auth/dto";

export async function signInAction(
  payload: SignInDto,
): Promise<AuthResponseDto> {
  const { data } = await api.post<AuthResponseDto>("/auth/login", payload);

  return data;
}

export async function signUpAction(
  payload: SignUpDto,
): Promise<AuthResponseDto> {
  const { data } = await api.post<AuthResponseDto>("/auth/register", payload);

  return data;
}
