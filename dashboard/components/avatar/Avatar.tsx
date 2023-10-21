import platform, { Integration, Provider } from '@utils/providerHelper';
import Image from 'next/image';

export type AvatarProps = {
  avatarName: Provider | Integration;
};

function Avatar({ avatarName }: AvatarProps) {
  const src = platform.getImgSrc(avatarName);
  return (
    <Image
      src={src}
      alt={`${avatarName} logo`}
      width={48}
      height={48}
      className="h-12 w-12 rounded-full border border-gray-100"
    />
  );
}

export default Avatar;
// !logo sizes are dynamic
