"use client";

import Link from "next/link";
import { AlertCircle, RefreshCw } from "lucide-react";
import { Button } from "@/shared/components/ui/button";
import ProfileHeader from "../components/ProfileHeader";
import ProfileInfoCard from "../components/ProfileInfoCard";
import ProfilePasswordCard from "../components/ProfilePasswordCard";
import { useProfile } from "../hooks/useProfile";

const Profile = () => {
  const {
    ready,
    loadError,
    reloadProfile,

    currentName,
    email,
    initials,

    name,
    setName,
    nameError,
    nameSuccess,
    saveName,

    currentPassword,
    setCurrentPassword,
    newPassword,
    setNewPassword,
    confirmPassword,
    setConfirmPassword,
    passwordError,
    passwordNotice,
    changePassword,
  } = useProfile();

  if (!ready) {
    return (
      <div className="flex w-full min-w-0 flex-col gap-8 px-5 py-7 sm:px-8 lg:px-10">
        <div className="bg-secondary-light h-12 w-48 animate-pulse rounded-xl" />

        <div className="bg-secondary-light h-44 animate-pulse rounded-3xl" />

        <div className="grid grid-cols-1 gap-6 xl:grid-cols-2">
          <div className="bg-secondary-light h-96 animate-pulse rounded-3xl" />
          <div className="bg-secondary-light h-96 animate-pulse rounded-3xl" />
        </div>
      </div>
    );
  }

  if (loadError) {
    return (
      <div className="flex min-h-[60vh] flex-col items-center justify-center gap-4 px-6 text-center">
        <div className="flex size-16 items-center justify-center rounded-2xl bg-red-50">
          <AlertCircle className="size-8 text-red-500" />
        </div>

        <h2 className="text-xl-bold text-primary-text">Gagal Memuat Profil</h2>

        <p className="text-sm-medium text-secondary-text max-w-md">
          {loadError}
        </p>

        <div className="flex flex-wrap gap-3">
          <Button
            type="button"
            onClick={() => void reloadProfile()}
            className="bg-primary hover:bg-secondary h-11 cursor-pointer rounded-xl px-5 text-white"
          >
            <RefreshCw className="size-4" />
            Coba Lagi
          </Button>

          <Button
            variant="outline"
            render={<Link href="/auth" />}
            className="h-11 rounded-xl px-5"
          >
            Kembali ke Login
          </Button>
        </div>
      </div>
    );
  }

  return (
    <div className="flex w-full min-w-0 flex-col gap-8 px-5 py-7 sm:px-8 lg:px-10">
      <header className="flex flex-col gap-2">
        <h1 className="text-3xl-bold text-primary-text">Profil Saya</h1>

        <p className="text-md-medium text-secondary-text">
          Kelola informasi pribadi dan keamanan akunmu.
        </p>
      </header>

      <ProfileHeader name={currentName} email={email} initials={initials} />

      <div className="grid grid-cols-1 items-start gap-6 xl:grid-cols-2">
        <ProfileInfoCard
          currentName={currentName}
          name={name}
          email={email}
          nameError={nameError}
          nameSuccess={nameSuccess}
          onNameChange={setName}
          onSave={saveName}
        />

        <ProfilePasswordCard
          currentPassword={currentPassword}
          newPassword={newPassword}
          confirmPassword={confirmPassword}
          passwordError={passwordError}
          passwordNotice={passwordNotice}
          onCurrentPasswordChange={setCurrentPassword}
          onNewPasswordChange={setNewPassword}
          onConfirmPasswordChange={setConfirmPassword}
          onSubmit={changePassword}
        />
      </div>
    </div>
  );
};

export default Profile;
