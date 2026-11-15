"use client";

import { useRouter } from "next/navigation";
import { useForm, type SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { isAxiosError } from "axios";
import type { SignFormValues, SignMode } from "../types/type";
import type { AuthResponseDto } from "@/shared/repository/auth/dto";
import { createSignSchema } from "@/shared/schemas/sign.schema";
import {
  useSignInMutation,
  useSignUpMutation,
} from "@/shared/repository/auth/query";

function getErrorMessage(error: unknown): string | null {
  if (!error) return null;

  if (isAxiosError<AuthResponseDto>(error)) {
    return (
      error.response?.data?.error ??
      error.response?.data?.message ??
      (error.response
        ? "Terjadi kesalahan. Silakan coba lagi."
        : "Gagal terhubung ke server.")
    );
  }

  return error instanceof Error
    ? error.message
    : "Terjadi kesalahan yang tidak diketahui.";
}

export function useSignForm(mode: SignMode) {
  const router = useRouter();

  const signInMutation = useSignInMutation();
  const signUpMutation = useSignUpMutation();

  const form = useForm<SignFormValues>({
    resolver: zodResolver(createSignSchema(mode)),
    defaultValues: {
      userName: "",
      email: "",
      password: "",
      confirmPassword: "",
    },
    mode: "onBlur",
  });

  const onSubmit: SubmitHandler<SignFormValues> = async (values) => {
    try {
      if (mode === "signIn") {
        await signInMutation.mutateAsync({
          email: values.email.trim(),
          password: values.password,
        });

        router.replace("/dashboard");
        router.refresh();
        return;
      }

      await signUpMutation.mutateAsync({
        userName: values.userName.trim(),
        email: values.email.trim(),
        password: values.password,
      });

      form.reset();
    } catch {}
  };

  const mutation = mode === "signIn" ? signInMutation : signUpMutation;

  return {
    form,
    onSubmit,
    isPending: form.formState.isSubmitting || mutation.isPending,
    error: getErrorMessage(mutation.error),
    success:
      mode === "signUp" && signUpMutation.isSuccess
        ? "Pendaftaran berhasil! Silakan masuk menggunakan akun baru."
        : null,
  };
}
