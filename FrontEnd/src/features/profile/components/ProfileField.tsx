"use client";

import { useState } from "react";
import type { ComponentProps } from "react";

import { Eye, EyeOff } from "lucide-react";

import { Input } from "@/shared/components/ui/input";
import { Button } from "@/shared/components/ui/button";
import {
  Field,
  FieldLabel,
  FieldDescription,
  FieldError,
} from "@/shared/components/ui/field";

interface ProfileFieldProps extends Omit<ComponentProps<typeof Input>, "id"> {
  id: string;
  label: string;
  helperText?: string;
  error?: string;
}

const ProfileField = ({
  id,
  label,
  helperText,
  error,
  type = "text",
  className = "",
  ...props
}: ProfileFieldProps) => {
  const [visible, setVisible] = useState(false);
  const isPassword = type === "password";

  return (
    <Field className="gap-2" data-invalid={Boolean(error)}>
      <FieldLabel htmlFor={id} className="text-sm-semibold text-primary-text">
        {label}
      </FieldLabel>

      <div className="relative">
        <Input
          {...props}
          id={id}
          type={isPassword && visible ? "text" : type}
          aria-invalid={Boolean(error)}
          className={`text-primary-text placeholder:text-secondary-text/60 focus-visible:border-secondary focus-visible:ring-secondary/10 read-only:bg-surface read-only:text-secondary-text h-12 w-full rounded-xl border border-[#DCE6ED] bg-white px-4 text-sm shadow-none transition-all duration-200 outline-none focus-visible:ring-3 ${isPassword ? "pr-12" : ""} ${className} `}
        />

        {isPassword && (
          <Button
            type="button"
            variant="ghost"
            size="icon"
            onClick={() => setVisible((prev) => !prev)}
            aria-label={visible ? "Sembunyikan password" : "Tampilkan password"}
            aria-pressed={visible}
            className="text-secondary-text hover:bg-secondary-light hover:text-primary absolute top-1/2 right-1 -translate-y-1/2 cursor-pointer rounded-lg"
          >
            {visible ? (
              <EyeOff className="size-4" />
            ) : (
              <Eye className="size-4" />
            )}
          </Button>
        )}
      </div>

      {helperText && (
        <FieldDescription className="text-xs">{helperText}</FieldDescription>
      )}

      {error && (
        <FieldError errors={[{ message: error }]} className="text-xs" />
      )}
    </Field>
  );
};

export default ProfileField;
