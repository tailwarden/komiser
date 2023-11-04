import platform, { IntegrationProvider, Provider } from '@utils/providerHelper';
import Image from 'next/image';

export type AvatarProps = {
  avatarName: Provider | IntegrationProvider;
  size?: number;
};

function Avatar({ avatarName, size = 24 }: AvatarProps) {
  const src = platform.getImgSrc(avatarName) || 'unknown platform';
  return (
    <Image
      src={src}
      alt={`${avatarName} logo`}
      width={size}
      height={size}
      className="rounded-full border border-gray-100"
    />
  );
}

export default Avatar;
