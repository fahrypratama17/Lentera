"use client";

import { ArrowRight, CheckCircle2, Loader2 } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { FieldGroup } from "@/shared/components/ui/field";
import AuthHeader from "./AuthHeader";
import AuthInput from "./AuthInput";
import { useSignForm } from "../hooks/useSignForm";
import type { RegisterFormProps } from "../types/type";

const RegisterForm = ({ onSwitch }: RegisterFormProps) => {
  const { form, onSubmit, isPending, error, success, resetFeedback } =
    useSignForm("signUp");

  const handleSwitch = () => {
    resetFeedback();
    onSwitch();
  };

  return (
    <div className="flex w-full flex-col gap-7">
      <AuthHeader
        title="Buat Akun Baru"
        description="Bergabung dan mulai belajar bersama."
      />

      <form
        onSubmit={form.handleSubmit(onSubmit)}
        noValidate
        className="flex flex-col gap-5"
      >
        <FieldGroup className="gap-4">
          <AuthInput
            control={form.control}
            name="userName"
            label="Nama Pengguna"
            placeholder="Masukkan nama pengguna"
            autoComplete="name"
          />

          <AuthInput
            control={form.control}
            name="email"
            label="Email"
            placeholder="Masukkan email"
            type="email"
            autoComplete="email"
          />

          <AuthInput
            control={form.control}
            name="password"
            label="Password"
            placeholder="Buat password"
            type="password"
            autoComplete="new-password"
          />

          <AuthInput
            control={form.control}
            name="confirmPassword"
            label="Konfirmasi Password"
            placeholder="Ulangi password"
            type="password"
            autoComplete="new-password"
          />
        </FieldGroup>

        {error && (
          <p
            role="alert"
            className="text-sm-medium rounded-xl bg-red-50 px-4 py-3 text-red-600"
          >
            {error}
          </p>
        )}

        {success && (
          <div
            role="status"
            className="border-secondary/15 bg-secondary-light flex flex-col gap-3 rounded-xl border p-4"
          >
            <div className="flex items-start gap-2">
              <CheckCircle2 className="text-secondary mt-0.5 size-5 shrink-0" />

              <p className="text-sm-medium text-primary">{success}</p>
            </div>

            <Button
              type="button"
              variant="outline"
              onClick={handleSwitch}
              className="border-secondary/20 text-primary h-10 cursor-pointer rounded-lg bg-white hover:bg-white/80"
            >
              Masuk Sekarang
              <ArrowRight className="size-4" />
            </Button>
          </div>
        )}

        <Button
          type="submit"
          size="lg"
          disabled={isPending}
          className="text-md-bold bg-primary hover:bg-secondary hover:shadow-secondary/15 h-12 w-full cursor-pointer rounded-xl text-white transition-all duration-200 hover:shadow-lg active:scale-[0.98] disabled:cursor-not-allowed disabled:opacity-60"
        >
          {isPending ? (
            <>
              <Loader2 className="size-4 animate-spin" />
              Memproses...
            </>
          ) : (
            <>
              Daftar Sekarang
              <ArrowRight className="size-4" />
            </>
          )}
        </Button>
      </form>

      <p className="text-sm-medium text-secondary-text text-center md:hidden">
        Sudah punya akun?{" "}
        <button
          type="button"
          onClick={handleSwitch}
          className="text-secondary cursor-pointer font-bold hover:underline"
        >
          Masuk
        </button>
      </p>
    </div>
  );
};

export default RegisterForm;
