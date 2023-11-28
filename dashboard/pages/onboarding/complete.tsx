import Head from 'next/head';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

import Button from '../../components/button/Button';
import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../components/onboarding-wizard/OnboardingWizardLayout';

export default function OnboardingComplete() {
  const router = useRouter();

  return (
    <div>
      <Head>
        <title>Onboarding Complete - Komiser</title>
        <meta name="description" content="Onboarding Complete - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout title="" progressBarWidth="1o0%">
          <div className="flex flex-col items-center justify-center p-12 pt-8 text-center text-gray-900/60">
            <Image
              src="/assets/img/purplin/rocket.svg"
              alt="Rocket"
              width={0}
              height={0}
              className="mb-8 h-56 w-56"
            />
            <h2 className="mb-3 text-2xl font-semibold leading-8 text-gray-950">
              Your data is being synced
            </h2>
            <div className="mb-10 font-normal">
              Processing time for your data varies based on its complexity. You
              can start using Komiser with the available data.
            </div>
            <Button
              onClick={() => router.push('/dashboard')}
              size="lg"
              style="primary"
              type="button"
            >
              Open Komiser
            </Button>
          </div>
        </LeftSideLayout>

        <RightSideLayout
          isCustom={true}
          customClasses="flex items-end justify-end"
        >
          <div className="relative">
            <Image
              src="/assets/img/others/empty-dashboard.png"
              alt="Dashboard"
              width={0}
              height={0}
              className="h-[600px] w-[500px]"
            />
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
