"use client";

import { useMutation } from "@tanstack/react-query";
import { signInAction, signUpAction } from "@/shared/repository/auth/action";
import type {
  SignInDto,
  SignUpDto,
  AuthResponseDto,
} from "@/shared/repository/auth/dto";

export function useSignInMutation() {
  return useMutation<AuthResponseDto, Error, SignInDto>({
    mutationKey: ["auth", "signIn"],
    mutationFn: signInAction,
    retry: false,
  });
}

export function useSignUpMutation() {
  return useMutation<AuthResponseDto, Error, SignUpDto>({
    mutationKey: ["auth", "signUp"],
    mutationFn: signUpAction,
    retry: false,
  });
}
