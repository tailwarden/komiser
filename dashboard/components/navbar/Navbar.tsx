import { useContext } from 'react';
import Image from 'next/image';
import Link from 'next/link';
import { useRouter } from 'next/router';

import useFeedbackWidget from '@components/feedback-widget/FeedbackWidget';
import GlobalAppContext from '../layout/context/GlobalAppContext';

interface NavItem {
  label: string;
  href: string;
}

function Navbar() {
  const { displayBanner, betaFlagOnboardingWizard } =
    useContext(GlobalAppContext);
  const router = useRouter();
  const { openFeedbackModal, FeedbackModal } = useFeedbackWidget();

  // TODO: (onboarding-wizard) Remove the betaFlagOnboardingWizard conditional when feature is stable
  const nav: NavItem[] = [
    { label: 'Dashboard', href: '/dashboard' },
    { label: 'Inventory', href: '/inventory' },
    betaFlagOnboardingWizard
      ? { label: 'Cloud Accounts', href: '/cloud-accounts' }
      : null,
    { label: 'Explorer', href: '/explorer' }
  ].filter(item => item !== null) as NavItem[];

  return (
    <nav
      className={`fixed ${
        displayBanner ? 'top-[72px]' : 'top-0'
      } z-30 flex w-full items-center justify-between gap-10 border-b border-gray-300 bg-white px-6 py-4 shadow-right xl:pr-8 2xl:pr-24`}
    >
      <div className="flex items-center gap-8 text-sm font-semibold text-gray-700">
        <Link href="/dashboard">
          <Image
            src="/assets/img/komiser.svg"
            width={40}
            height={40}
            alt="Komiser logo"
          />
        </Link>
        {nav.map((navItem, idx) => (
          <Link
            key={idx}
            href={navItem.href}
            className={
              router.pathname === navItem.href
                ? 'text-darkcyan-500'
                : 'text-gray-700'
            }
          >
            {navItem.label}
          </Link>
        ))}
      </div>
      <div className="flex gap-4 text-sm font-medium text-gray-950 lg:gap-10">
        <a
          className="hidden items-center gap-2 transition-colors hover:text-darkcyan-500 md:flex"
          href="https://docs.komiser.io/docs/intro?utm_source=komiser&utm_medium=referral&utm_campaign=static"
          target="_blank"
          rel="noopener noreferrer"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeMiterlimit="10"
              strokeWidth="2"
              d="M21 7v10c0 3-1.5 5-5 5H8c-3.5 0-5-2-5-5V7c0-3 1.5-5 5-5h8c3.5 0 5 2 5 5z"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeMiterlimit="10"
              strokeWidth="2"
              d="M14.5 4.5v2c0 1.1.9 2 2 2h2M10 13l-2 2 2 2M14 13l2 2-2 2"
            ></path>
          </svg>
          Docs
        </a>
        <a
          className="hidden items-center gap-2 transition-colors hover:text-darkcyan-500 md:flex"
          href="https://www.tailwarden.com/changelog/komiser"
          target="_blank"
          rel="noopener noreferrer"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeMiterlimit="10"
              strokeWidth="2"
              d="M8 2v3M16 2v3M7 11h8M7 15h5M15 22H9c-5 0-6-2.06-6-6.18V9.65c0-4.7 1.67-5.96 5-6.15h8c3.33.18 5 1.45 5 6.15V16"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M21 16l-6 6v-3c0-2 1-3 3-3h3z"
            ></path>
          </svg>
          Changelog
        </a>
        <a
          className="hidden cursor-pointer items-center gap-2 transition-colors hover:text-darkcyan-500 md:flex"
          rel="noopener noreferrer"
          onClick={() => openFeedbackModal()}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="16"
            height="16"
            fill="none"
            viewBox="0 0 24 25"
          >
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth="2"
              d="M11 2.75H9c-5 0-7 2-7 7v6c0 5 2 7 7 7h6c5 0 7-2 7-7v-2"
            ></path>
            <path
              stroke="currentColor"
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeMiterlimit="10"
              strokeWidth="2"
              d="M16.04 3.77l-7.88 7.88c-.3.3-.6.89-.66 1.32l-.43 3.01c-.16 1.09.61 1.85 1.7 1.7l3.01-.43c.42-.06 1.01-.36 1.32-.66l7.88-7.88c1.36-1.36 2-2.94 0-4.94-2-2-3.58-1.36-4.94 0zM14.91 4.9a7.144 7.144 0 004.94 4.94"
            ></path>
          </svg>
          Leave feedback
        </a>
        <a
          className="flex items-center gap-2 rounded-lg bg-[#5865F2] px-4 py-2 text-white transition-colors hover:bg-[#4f5be2]"
          href="https://discord.tailwarden.com"
          target="_blank"
          rel="noopener noreferrer"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            fill="none"
            viewBox="0 0 24 24"
          >
            <path
              fill="currentColor"
              d="M18.93 4.935a16.457 16.457 0 00-4.07-1.266.062.062 0 00-.066.031c-.175.314-.37.723-.506 1.044a15.183 15.183 0 00-4.573 0c-.136-.328-.338-.73-.515-1.044a.064.064 0 00-.065-.031 16.413 16.413 0 00-4.07 1.266.058.058 0 00-.028.023c-2.593 3.885-3.303 7.674-2.954 11.417a.069.069 0 00.026.047 16.565 16.565 0 004.994 2.531.065.065 0 00.07-.023c.385-.527.728-1.082 1.022-1.666a.064.064 0 00-.035-.089 10.906 10.906 0 01-1.56-.745.064.064 0 01-.007-.107c.105-.079.21-.16.31-.244a.061.061 0 01.065-.008c3.273 1.498 6.817 1.498 10.051 0a.062.062 0 01.066.008c.1.082.204.165.31.244a.064.064 0 01-.005.107c-.499.292-1.017.538-1.561.744a.064.064 0 00-.034.09c.3.583.643 1.139 1.02 1.666a.063.063 0 00.07.023 16.51 16.51 0 005.003-2.531.065.065 0 00.026-.047c.417-4.326-.699-8.084-2.957-11.416a.05.05 0 00-.026-.024zM8.684 14.096c-.985 0-1.797-.907-1.797-2.022 0-1.114.796-2.021 1.797-2.021 1.01 0 1.813.915 1.798 2.021 0 1.115-.796 2.022-1.798 2.022zm6.646 0c-.986 0-1.797-.907-1.797-2.022 0-1.114.796-2.021 1.797-2.021 1.009 0 1.813.915 1.797 2.021 0 1.115-.788 2.022-1.797 2.022z"
            ></path>
          </svg>
          Join our Discord
        </a>
      </div>
      <FeedbackModal />
    </nav>
  );
}

export default Navbar;
