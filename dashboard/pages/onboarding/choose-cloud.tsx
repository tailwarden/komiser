import { useState } from 'react';
import Head from 'next/head';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

import Avatar from '@components/avatar/Avatar';
import platform, { allProviders, Provider } from '../../utils/providerHelper';

import Button from '../../components/button/Button';
import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../components/onboarding-wizard/OnboardingWizardLayout';
import SelectInput from '../../components/onboarding-wizard/SelectInput';

export default function ChooseCloud() {
  const router = useRouter();
  const [provider, setProvider] = useState<Provider>(allProviders.AWS);

  const handleNext = () => {
    router.push(`/onboarding/provider/${provider}`);
  };

  const handleSuggest = () =>
    router.replace(
      'https://docs.komiser.io/docs/faqs#how-can-i-request-a-new-feature'
    );

  const handleSelectChange = (newValue: string) =>
    setProvider(newValue as Provider);

  return (
    <div>
      <Head>
        <title>Onboarding - Komiser</title>
        <meta name="description" content="Onboarding - Komiser" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <OnboardingWizardLayout>
        <LeftSideLayout
          title="Connect your cloud account"
          progressBarWidth="35%"
        >
          <div className="leading-6 text-gray-900/60">
            <div className="font-normal">
              Komiser is cloud agnostic, one platform across all major public
              cloud <br /> providers.
            </div>
            <div>Get started now by connecting your first account.</div>
          </div>
          <div className="py-10">
            <SelectInput
              label="Cloud provider"
              value={provider}
              values={Object.values(allProviders)}
              handleChange={handleSelectChange}
              displayValues={Object.values(allProviders).map(value => ({
                label: platform.getLabel(value)
              }))}
            />
          </div>
          <div className="flex justify-between">
            <Button
              onClick={handleSuggest}
              size="lg"
              style="text"
              type="button"
            >
              Suggest a cloud provider
            </Button>
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

        <RightSideLayout
          isCustom={true}
          customClasses="flex flex-col justify-center items-center space-y-6"
        >
          <div className="relative inline-block">
            <Image
              src="/assets/img/others/onboarding-cloud.svg"
              alt="Onboarding cloud"
              width={500}
              height={120}
            />
            <div className="absolute left-[48%] top-[53%] -translate-x-1/2 -translate-y-1/2 transform rounded-full">
              <Avatar avatarName={provider} size={96} />
            </div>
          </div>
          <Image
            width={20}
            height={20}
            alt="Arrow down"
            src="/assets/img/others/arrow-down.svg"
          />
          <Image
            alt={`${provider} Logo`}
            src={'/assets/img/komiser.svg'}
            width={120}
            height={120}
          />
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
