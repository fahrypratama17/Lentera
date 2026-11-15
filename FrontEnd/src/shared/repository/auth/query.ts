"use client";

import { useMutation } from "@tanstack/react-query";
import { signInAction, signUpAction } from "./action";

import type {
  SignInDto,
  SignUpDto,
  SignInResponseDto,
  AuthResponseDto,
} from "./dto";

export function useSignInMutation() {
  return useMutation<SignInResponseDto, Error, SignInDto>({
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
