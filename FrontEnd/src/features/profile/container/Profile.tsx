"use client";

import ProfileHeader from "../components/ProfileHeader";
import ProfileInfoCard from "../components/ProfileInfoCard";
import ProfilePasswordCard from "../components/ProfilePasswordCard";

import { useProfile } from "../hooks/useProfile";

const Profile = () => {
  const {
    ready,
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
