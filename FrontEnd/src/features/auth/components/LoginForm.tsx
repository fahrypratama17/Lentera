"use client";

import { ArrowRight, Loader2 } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { FieldGroup } from "@/shared/components/ui/field";
import AuthHeader from "./AuthHeader";
import AuthInput from "./AuthInput";
import { useSignForm } from "@/features/auth/hooks/useSignForm";
import { LoginFormProps } from "@/features/auth/types/type";

const LoginForm = ({ onSwitch }: LoginFormProps) => {
  const { form, onSubmit, isPending, error } = useSignForm("signIn");

  return (
    <div className="flex w-full flex-col gap-8">
      <AuthHeader
        title="Selamat Datang!"
        description="Masuk dan lanjutkan perjalanan belajarmu."
      />

      <form
        onSubmit={form.handleSubmit(onSubmit)}
        noValidate
        className="flex flex-col gap-6"
      >
        <FieldGroup className="gap-5">
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
            placeholder="Masukkan password"
            type="password"
            autoComplete="current-password"
          />
        </FieldGroup>

        {error && (
          <p role="alert" className="text-sm-medium text-red-500">
            {error}
          </p>
        )}

        <Button
          type="submit"
          size="lg"
          disabled={isPending}
          className="text-md-bold bg-primary hover:bg-secondary hover:shadow-secondary/15 h-12 w-full cursor-pointer rounded-xl text-white transition-all duration-200 hover:shadow-lg active:scale-[0.98]"
        >
          {isPending ? (
            <>
              <Loader2 className="size-4 animate-spin" />
              Memproses...
            </>
          ) : (
            <>
              Masuk
              <ArrowRight className="size-4" />
            </>
          )}
        </Button>
      </form>

      <p className="text-sm-medium text-secondary-text text-center md:hidden">
        Belum punya akun?{" "}
        <button
          type="button"
          onClick={onSwitch}
          className="text-secondary cursor-pointer font-bold hover:underline"
        >
          Daftar sekarang
        </button>
      </p>
    </div>
  );
};

export default LoginForm;
