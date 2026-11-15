"use client";

import type { FormEventHandler } from "react";
import { Save, UserRound } from "lucide-react";
import { Card, CardContent } from "@/shared/components/ui/card";
import { Button } from "@/shared/components/ui/button";
import ProfileField from "./ProfileField";

interface ProfileInfoCardProps {
  currentName: string;
  name: string;
  email: string;
  nameError: string;
  nameSuccess: string;
  onNameChange: (value: string) => void;
  onSave: FormEventHandler<HTMLFormElement>;
}

const ProfileInfoCard = ({
  currentName,
  name,
  email,
  nameError,
  nameSuccess,
  onNameChange,
  onSave,
}: ProfileInfoCardProps) => {
  const isUnchanged = name.trim() === currentName;

  return (
    <Card className="gap-0 rounded-3xl border border-[#DCE6ED] bg-white py-0 shadow-sm">
      <CardContent className="gap-0 px-6 py-6 md:px-8 md:py-8">
        <div className="mb-7 flex items-center gap-3">
          <div className="bg-secondary-light flex size-12 shrink-0 items-center justify-center rounded-2xl">
            <UserRound size={23} className="text-secondary" />
          </div>

          <div>
            <h2 className="text-xl-bold text-primary-text">
              Informasi Pribadi
            </h2>
            <p className="text-sm-medium text-secondary-text mt-1">
              Perbarui nama yang ditampilkan.
            </p>
          </div>
        </div>

        <form onSubmit={onSave} noValidate className="flex flex-col gap-5">
          <ProfileField
            id="profile-name"
            label="Nama Lengkap"
            type="text"
            value={name}
            onChange={(event) => onNameChange(event.target.value)}
            placeholder="Masukkan nama lengkap"
            autoComplete="name"
            minLength={2}
            maxLength={100}
            required
          />

          <ProfileField
            id="profile-email"
            label="Alamat Email"
            type="email"
            value={email}
            readOnly
            helperText="Email tidak dapat diubah pada halaman ini."
          />

          {nameError && (
            <p
              role="alert"
              className="rounded-xl bg-red-50 px-4 py-3 text-sm text-red-600"
            >
              {nameError}
            </p>
          )}

          {nameSuccess && (
            <p
              role="status"
              className="bg-secondary-light text-primary rounded-xl px-4 py-3 text-sm"
            >
              {nameSuccess}
            </p>
          )}

          <div className="mt-3 border-t border-[#DCE6ED] pt-5">
            <Button
              type="submit"
              size="lg"
              disabled={isUnchanged}
              className="text-sm-semibold bg-primary hover:bg-secondary h-11 cursor-pointer rounded-xl px-5 text-white transition-colors disabled:cursor-not-allowed disabled:opacity-50"
            >
              <Save className="size-4" />
              Simpan Perubahan
            </Button>
          </div>
        </form>
      </CardContent>
    </Card>
  );
};

export default ProfileInfoCard;
