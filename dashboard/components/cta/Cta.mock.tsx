import { CtaProps } from './Cta';

const base: CtaProps = {
  title: 'Introducing Tailwarden',
  description:
    'Tailwarden is the cloud version of Komiser, which offers more features and insights',
  action: (
    <button className="flex items-center gap-2 rounded-lg bg-purple-500 px-4 py-2 text-white transition-colors hover:bg-purple-600">
      <svg
        width="18"
        height="15"
        viewBox="0 0 18 15"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          d="M0.875 10.1319C0.875 12.203 2.55393 13.8819 4.625 13.8819H14C15.7259 13.8819 17.125 12.4828 17.125 10.7569C17.125 9.42181 16.2878 8.28228 15.1097 7.83466C15.2006 7.57559 15.25 7.29701 15.25 7.0069C15.25 5.62619 14.1307 4.5069 12.75 4.5069C12.4806 4.5069 12.2212 4.54951 11.9781 4.62835C11.4804 2.75906 9.77602 1.3819 7.75 1.3819C5.33375 1.3819 3.375 3.34065 3.375 5.7569C3.375 6.03458 3.40087 6.30623 3.45033 6.56956C1.95463 7.06248 0.875 8.47111 0.875 10.1319Z"
          stroke="white"
          stroke-width="1.25"
          stroke-linecap="round"
          stroke-linejoin="round"
        />
      </svg>
      Discover Tailwarden
    </button>
  )
};

const mockCTaProps = {
  base
};

export default mockCTaProps;
