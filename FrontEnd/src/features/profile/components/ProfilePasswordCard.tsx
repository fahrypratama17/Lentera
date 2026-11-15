"use client";

import type { FormEventHandler } from "react";
import { KeyRound, Save } from "lucide-react";
import { Card, CardContent } from "@/shared/components/ui/card";
import { Button } from "@/shared/components/ui/button";
import ProfileField from "./ProfileField";

interface ProfilePasswordCardProps {
  currentPassword: string;
  newPassword: string;
  confirmPassword: string;
  passwordError: string;
  passwordNotice: string;
  onCurrentPasswordChange: (value: string) => void;
  onNewPasswordChange: (value: string) => void;
  onConfirmPasswordChange: (value: string) => void;
  onSubmit: FormEventHandler<HTMLFormElement>;
}

const ProfilePasswordCard = ({
  currentPassword,
  newPassword,
  confirmPassword,
  passwordError,
  passwordNotice,
  onCurrentPasswordChange,
  onNewPasswordChange,
  onConfirmPasswordChange,
  onSubmit,
}: ProfilePasswordCardProps) => {
  return (
    <Card className="gap-0 rounded-3xl border border-[#DCE6ED] bg-white py-0 shadow-sm">
      <CardContent className="gap-0 px-6 py-6 md:px-8 md:py-8">
        <div className="mb-7 flex items-center gap-3">
          <div className="bg-secondary-light flex size-12 shrink-0 items-center justify-center rounded-2xl">
            <KeyRound size={23} className="text-secondary" />
          </div>

          <div>
            <h2 className="text-xl-bold text-primary-text">Ubah Password</h2>
            <p className="text-sm-medium text-secondary-text mt-1">
              Kelola keamanan akunmu.
            </p>
          </div>
        </div>

        <form onSubmit={onSubmit} noValidate className="flex flex-col gap-5">
          <ProfileField
            id="current-password"
            label="Password Saat Ini"
            type="password"
            value={currentPassword}
            onChange={(event) => onCurrentPasswordChange(event.target.value)}
            placeholder="Masukkan password lama"
            autoComplete="current-password"
            required
          />

          <ProfileField
            id="new-password"
            label="Password Baru"
            type="password"
            value={newPassword}
            onChange={(event) => onNewPasswordChange(event.target.value)}
            placeholder="Masukkan password baru"
            autoComplete="new-password"
            minLength={8}
            helperText="Gunakan minimal 8 karakter."
            required
          />

          <ProfileField
            id="confirm-password"
            label="Konfirmasi Password Baru"
            type="password"
            value={confirmPassword}
            onChange={(event) => onConfirmPasswordChange(event.target.value)}
            placeholder="Ulangi password baru"
            autoComplete="new-password"
            required
          />

          {passwordError && (
            <p
              role="alert"
              className="rounded-xl bg-red-50 px-4 py-3 text-sm text-red-600"
            >
              {passwordError}
            </p>
          )}

          {passwordNotice && (
            <p
              role="status"
              className="bg-secondary-light text-primary rounded-xl px-4 py-3 text-sm"
            >
              {passwordNotice}
            </p>
          )}

          <div className="mt-3 border-t border-[#DCE6ED] pt-5">
            <Button
              type="submit"
              size="lg"
              className="text-sm-semibold bg-primary hover:bg-secondary h-11 cursor-pointer rounded-xl px-5 text-white transition-colors"
            >
              <Save className="size-4" />
              Validasi Password
            </Button>
          </div>
        </form>
      </CardContent>
    </Card>
  );
};

export default ProfilePasswordCard;
