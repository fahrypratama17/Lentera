"use client";

import { useMutation } from "@tanstack/react-query";
import { signInAction, signOutAction, signUpAction } from "./action";

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

export function useSignOutMutation() {
  return useMutation<void, Error>({
    mutationKey: ["auth", "signOut"],
    mutationFn: signOutAction,
    retry: false,
  });
}
