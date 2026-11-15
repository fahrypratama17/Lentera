"use client";

import { useRouter } from "next/navigation";
import { useForm, type SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { isAxiosError } from "axios";
import type { SignFormValues, SignMode } from "../types/type";
import type { AuthErrorDto } from "@/shared/repository/auth/dto";
import { createSignSchema } from "@/shared/schemas/sign.schema";
import {
  useSignInMutation,
  useSignUpMutation,
} from "@/shared/repository/auth/query";

function getErrorMessage(error: unknown): string | null {
  if (!error) return null;

  if (isAxiosError<AuthErrorDto>(error)) {
    if (!error.response) {
      return "Gagal terhubung ke server.";
    }

    return (
      error.response.data?.error ??
      error.response.data?.message ??
      "Terjadi kesalahan. Silakan coba lagi."
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

  const onSubmit: SubmitHandler<SignFormValues> = (values) => {
    if (mode === "signIn") {
      if (signInMutation.isPending) return;

      signInMutation.reset();

      signInMutation.mutate(
        {
          email: values.email.trim(),
          password: values.password,
        },
        {
          onSuccess: (response) => {
            const token = response.data?.access_token;

            if (!token) {
              form.setError("root", {
                message: "Access token tidak ditemukan.",
              });
              return;
            }

            localStorage.setItem("lentera_token", token);

            router.replace("/dashboard");
            router.refresh();
          },
        },
      );

      return;
    }

    if (signUpMutation.isPending) return;

    signUpMutation.reset();

    signUpMutation.mutate(
      {
        Name: values.userName.trim(),
        Email: values.email.trim(),
        Password: values.password,
      },
      {
        onSuccess: () => {
          form.reset();
        },
      },
    );
  };

  const mutation = mode === "signIn" ? signInMutation : signUpMutation;

  const resetFeedback = () => {
    mutation.reset();
  };

  return {
    form,
    onSubmit,
    isPending: mutation.isPending,
    error: getErrorMessage(mutation.error),
    success:
      mode === "signUp" && signUpMutation.isSuccess
        ? "Pendaftaran berhasil! Silakan masuk menggunakan akun baru."
        : null,
    resetFeedback,
  };
}
