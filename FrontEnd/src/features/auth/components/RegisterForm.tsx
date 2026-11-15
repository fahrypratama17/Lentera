import { ArrowRight, Loader2 } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import { FieldGroup } from "@/shared/components/ui/field";
import AuthHeader from "./AuthHeader";
import AuthInput from "./AuthInput";
import { useSignForm } from "../hooks/useSignForm";

interface RegisterFormProps {
  onSwitch: () => void;
}

const RegisterForm = ({ onSwitch }: RegisterFormProps) => {
  const { form, onSubmit, isPending, error } = useSignForm("signUp");

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
            label="Username"
            placeholder="Masukkan username"
            autoComplete="username"
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
          onClick={onSwitch}
          className="text-secondary cursor-pointer font-bold hover:underline"
        >
          Masuk
        </button>
      </p>
    </div>
  );
};

export default RegisterForm;
