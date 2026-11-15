"use client";

import { useState } from "react";
import { useRouter } from "next/navigation";
import { useForm, type SubmitHandler } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import type { ApiResponse, SignFormValues, SignMode } from "../types/type";
import { createSignSchema } from "@/shared/schemas/sign.schema";

const API_URL = {
  signIn: "/api/auth/login",
  signUp: "/api/auth/register",
};

export function useSignForm(mode: SignMode) {
  const router = useRouter();

  const [error, setError] = useState<string | null>(null);
  const [success, setSuccess] = useState<string | null>(null);

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
    setError(null);
    setSuccess(null);

    try {
      const payload =
        mode === "signIn"
          ? {
              email: values.email.trim(),
              password: values.password,
            }
          : {
              userName: values.userName.trim(),
              email: values.email.trim(),
              password: values.password,
            };

      const response = await fetch(API_URL[mode], {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        credentials: "same-origin",
        body: JSON.stringify(payload),
      });

      const data = (await response
        .json()
        .catch(() => null)) as ApiResponse | null;

      if (!response.ok) {
        throw new Error(
          data?.error ??
            data?.message ??
            "Terjadi kesalahan. Silakan coba lagi.",
        );
      }

      if (mode === "signIn") {
        router.replace("/dashboard");
        router.refresh();
        return;
      }

      setSuccess("Pendaftaran berhasil! Silakan masuk menggunakan akun baru.");

      form.reset();
    } catch (err) {
      setError(
        err instanceof Error
          ? err.message
          : "Terjadi kesalahan yang tidak diketahui.",
      );
    }
  };

  return {
    form,
    onSubmit,
    isPending: form.formState.isSubmitting,
    error,
    success,
  };
}
