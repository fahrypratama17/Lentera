import { api } from "@/shared/lib/axios";

import type {
  SignInDto,
  SignUpDto,
  SignInResponseDto,
  AuthResponseDto,
} from "./dto";

export async function signInAction(
  payload: SignInDto,
): Promise<SignInResponseDto> {
  const { data } = await api.post<SignInResponseDto>("/auth/login", payload);

  return data;
}

export async function signUpAction(
  payload: SignUpDto,
): Promise<AuthResponseDto> {
  const { data } = await api.post<AuthResponseDto>("/auth/register", payload);

  return data;
}
