import { useState } from 'react';
import Head from 'next/head';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

import { DBProvider, allDBProviders } from '../../utils/providerHelper';

import Button from '../../components/button/Button';
import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../components/onboarding-wizard/OnboardingWizardLayout';

interface DatabaseItemProps {
  value: DBProvider;
  selected: DBProvider;
  label?: string;
  imageUrl: string;
  handleClick: (db: DBProvider) => void;
}

function DatabaseLeftItem({
  imageUrl,
  label,
  value,
  selected,
  handleClick
}: DatabaseItemProps) {
  const onClick = () => handleClick(value);

  return (
    <div
      onClick={onClick}
      className={`flex cursor-pointer flex-col items-center justify-center rounded-lg border-[1.5px] p-6 ${
        selected === value
          ? 'border-darkcyan-500'
          : 'border-gray-200 hover:border-cyan-400'
      }`}
    >
      <Image
        src={imageUrl}
        alt={`${label} Logo`}
        className="h-14 w-14 rounded-full"
        width={0}
        height={0}
      />
      <div>{label}</div>
    </div>
  );
}

function DatabaseRightItem({
  imageUrl,
  value,
  selected,
  handleClick
}: DatabaseItemProps) {
  const onClick = () => handleClick(value);

  return (
    <div
      onClick={onClick}
      key={value}
      className={`flex w-32 items-center justify-center rounded-3xl p-6 ${
        selected === value ? 'bg-cyan-200' : 'bg-white'
      }`}
    >
      <Image
        src={imageUrl}
        alt={`${value} Logo`}
        className={`h-20 w-20 rounded-full ${
          selected !== value ? 'opacity-[0.6]' : ''
        }`}
        width={0}
        height={0}
      />
    </div>
  );
}

export default function ChooseDatabase() {
  const router = useRouter();
  const [database, setDatabase] = useState<DBProvider>(allDBProviders.POSTGRES);

  const handleNext = () => router.push(`/onboarding/database/${database}`);

  const handleClick = (db: DBProvider) => setDatabase(db);

  return (
    <div>
      <Head>
        <title>Select Database - Komiser</title>
        <meta name="description" content="Select Database - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Persist your account data"
          progressBarWidth="70%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              Add a way to store and retain data through a database, so that it
              remains accessible and preserved even after you end your session
              on Komiser.
            </div>
          </div>
          <div className="py-10">
            <p className="pb-8">Select a database type</p>
            <div className="grid grid-cols-2 gap-4">
              <DatabaseLeftItem
                imageUrl="/assets/img/database/postgresql.svg"
                label="PostgreSQL"
                value={allDBProviders.POSTGRES}
                selected={database}
                handleClick={handleClick}
              />
              <DatabaseLeftItem
                imageUrl="/assets/img/database/sqlite.svg"
                label="SQLite"
                value={allDBProviders.SQLITE}
                selected={database}
                handleClick={handleClick}
              />
            </div>
          </div>
          <div className="flex flex-row-reverse">
            <Button
              onClick={handleNext}
              size="lg"
              style="primary"
              type="button"
            >
              Next
            </Button>
          </div>
        </LeftSideLayout>

        <RightSideLayout isCustom={true} customClasses="flex justify-center">
          <div className="relative">
            <Image
              src="/assets/img/database/selectdb-komiser.svg"
              alt="Komiser Logo"
              width={400}
              height={150}
              className="-mt-0"
            />

            <div className="mt-10">
              <Image
                src="/assets/img/database/db-arrow-vector.svg"
                alt="Komiser Logo"
                width={135}
                height={100}
                className={`${
                  database === allDBProviders.POSTGRES
                    ? 'ml-16'
                    : 'rotate-onboarding-arrow ml-[182px]'
                }`}
              />
            </div>

            <div className="-ml-3 mt-10 grid grid-cols-2 justify-items-center gap-[90px]">
              <DatabaseRightItem
                imageUrl="/assets/img/database/postgresql.svg"
                value={allDBProviders.POSTGRES}
                selected={database}
                handleClick={handleClick}
              />
              <DatabaseRightItem
                imageUrl="/assets/img/database/sqlite.svg"
                value={allDBProviders.SQLITE}
                selected={database}
                handleClick={handleClick}
              />
            </div>
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
