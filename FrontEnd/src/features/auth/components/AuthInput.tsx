"use client";

import { useId, useState } from "react";
import {
  Controller,
  type Control,
  type FieldValues,
  type Path,
} from "react-hook-form";
import { Eye, EyeOff } from "lucide-react";

import { Field, FieldLabel, FieldError } from "@/shared/components/ui/field";
import { Button } from "@/shared/components/ui/button";

interface AuthInputProps<T extends FieldValues> {
  control: Control<T>;
  name: Path<T>;
  label: string;
  placeholder: string;
  type?: "text" | "email" | "password";
  autoComplete?: string;
}

export default function AuthInput<T extends FieldValues>({
  control,
  name,
  label,
  placeholder,
  type = "text",
  autoComplete,
}: AuthInputProps<T>) {
  const id = useId();
  const errorId = `${id}-error`;

  const [showPassword, setShowPassword] = useState(false);

  const isPassword = type === "password";

  return (
    <Controller
      control={control}
      name={name}
      render={({ field, fieldState }) => (
        <Field data-invalid={fieldState.invalid} className="gap-2">
          <FieldLabel
            htmlFor={id}
            className="text-sm-semibold text-primary-text"
          >
            {label}
          </FieldLabel>

          <div className="relative">
            <input
              {...field}
              id={id}
              value={typeof field.value === "string" ? field.value : ""}
              type={isPassword && showPassword ? "text" : type}
              placeholder={placeholder}
              autoComplete={autoComplete}
              aria-invalid={fieldState.invalid}
              aria-describedby={fieldState.invalid ? errorId : undefined}
              className={`text-sm-medium border-primary/15 bg-surface text-primary-text placeholder:text-secondary-text/60 focus:border-secondary focus:ring-secondary/10 h-12 w-full rounded-xl border px-4 transition-all duration-200 outline-none focus:ring-4 aria-invalid:border-red-400 ${
                isPassword ? "pr-12" : "pr-4"
              } `}
            />

            {isPassword && (
              <Button
                type="button"
                variant="ghost"
                size="icon"
                onClick={() => setShowPassword((prev) => !prev)}
                aria-label={
                  showPassword ? "Sembunyikan password" : "Tampilkan password"
                }
                aria-pressed={showPassword}
                className="text-secondary-text hover:bg-secondary-light hover:text-primary absolute top-1/2 right-1 -translate-y-1/2 rounded-lg"
              >
                {showPassword ? <EyeOff size={18} /> : <Eye size={18} />}
              </Button>
            )}
          </div>

          <FieldError
            id={errorId}
            errors={[fieldState.error]}
            className="text-xs"
          />
        </Field>
      )}
    />
  );
}
