import { useState } from 'react';
import Head from 'next/head';
import Image from 'next/image';
import { useRouter } from 'next/navigation';

import Select from '../../components/select/Select';
import Button from '../../components/button/Button';
import OnboardingWizardLayout, {
  LeftSideLayout,
  RightSideLayout
} from '../../components/onboarding-wizard/OnboardingWizardLayout';

const SelectCloud = {
  'Amazon Web Services': 'aws',
  'Microsoft Azure': 'azure',
  'Google Cloud Platform': 'gcp',
  'Oracle Cloud Infrastructure': 'oci',
  Kubernetes: 'kubernetes',
  'Digital Ocean': 'digitalocean',
  Civo: 'civo',
  'MongoDB Atlas': 'mongodbatlas',
  'Tencent Cloud': 'tencent',
  Scaleway: 'scaleway',
  'OVH Cloud': 'ovh',
  Linode: 'linode'
} as const;

type Clouds = keyof typeof SelectCloud;
type CloudsValues = (typeof SelectCloud)[Clouds];

const cloudLogo: { [K in CloudsValues]: string } = {
  aws: '/assets/img/providers/aws.png',
  azure: '/assets/img/providers/azure.svg',
  gcp: '/assets/img/providers/gcp.png',
  oci: '/assets/img/providers/oci.png',
  civo: '/assets/img/providers/civo.jpeg',
  tencent: '/assets/img/providers/tencent.jpeg',
  kubernetes: '/assets/img/providers/kubernetes.png',
  digitalocean: '/assets/img/providers/digitalocean.png',
  mongodbatlas: '/assets/img/providers/mongodbatlas.jpg',
  scaleway: '/assets/img/providers/scaleway.png',
  ovh: '/assets/img/providers/ovh.jpeg',
  linode: '/assets/img/providers/linode.png'
};

export default function Onboarding() {
  const router = useRouter();
  const [provider, setProvider] = useState<CloudsValues>('aws');

  const handleNext = () => {
    router.push(`/onboarding/${provider}`);
  };

  const handleSuggest = () =>
    router.replace(
      'https://docs.komiser.io/docs/faqs#how-can-i-request-a-new-feature'
    );

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
            <Select
              displayValues={Object.keys(SelectCloud)}
              handleChange={(value: any) => setProvider(value)}
              label="Cloud provider"
              value={provider}
              values={Object.values(SelectCloud)}
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

        <RightSideLayout>
          <div className="relative">
            <Image
              src="/assets/img/others/onboarding-padlock.png"
              alt="Komiser Logo"
              width={500}
              height={150}
            />
            <div className="absolute top-0 left-0 mt-[10.1rem] ml-[14.55rem] flex h-[75px] w-[77px] items-center justify-center rounded-full border bg-gray-800 p-[0.1rem]">
              <Image
                src={cloudLogo[provider]}
                alt={`${provider} Logo`}
                className="h-full w-full rounded-full"
                width={0}
                height={0}
              />
            </div>
          </div>
        </RightSideLayout>
      </OnboardingWizardLayout>
    </div>
  );
}
