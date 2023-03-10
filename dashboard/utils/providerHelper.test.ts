import providers from './providerHelper';

describe('providerHelper function outputs', () => {
  it('should return the right label for aws', () => {
    const result = providers.providerLabel('aws');
    expect(result).toBe('Amazon Web Services');
  });

  it('should return the right label for gcp', () => {
    const result = providers.providerLabel('gcp');
    expect(result).toBe('Google Cloud Platform');
  });

  it('should return the right label for ovh', () => {
    const result = providers.providerLabel('ovh');
    expect(result).toBe('OVH');
  });

  it('should return the right label for digitalocean', () => {
    const result = providers.providerLabel('digitalocean');
    expect(result).toBe('DigitalOcean');
  });

  it('should return the right label for azure', () => {
    const result = providers.providerLabel('azure');
    expect(result).toBe('Azure');
  });

  it('should return the right label for civo', () => {
    const result = providers.providerLabel('civo');
    expect(result).toBe('Civo');
  });

  it('should return the right label for kubernetes', () => {
    const result = providers.providerLabel('kubernetes');
    expect(result).toBe('Kubernetes');
  });

  it('should return the right label for linode', () => {
    const result = providers.providerLabel('linode');
    expect(result).toBe('Linode');
  });

  it('should return the right label for tencent', () => {
    const result = providers.providerLabel('tencent');
    expect(result).toBe('Tencent');
  });

  it('should return the right label for oci', () => {
    const result = providers.providerLabel('oci');
    expect(result).toBe('OCI');
  });

  it('should return the right label for scaleway', () => {
    const result = providers.providerLabel('scaleway');
    expect(result).toBe('Scaleway');
  });

  it('should return the right img for aws', () => {
    const result = providers.providerImg('aws');
    expect(result).toBe('/assets/img/providers/aws.png');
  });

  it('should return the right img for gcp', () => {
    const result = providers.providerImg('gcp');
    expect(result).toBe('/assets/img/providers/gcp.png');
  });

  it('should return the right img for ovh', () => {
    const result = providers.providerImg('ovh');
    expect(result).toBe('/assets/img/providers/ovh.jpeg');
  });

  it('should return the right img for digitalocean', () => {
    const result = providers.providerImg('digitalocean');
    expect(result).toBe('/assets/img/providers/digitalocean.png');
  });

  it('should return the right img for azure', () => {
    const result = providers.providerImg('azure');
    expect(result).toBe('/assets/img/providers/azure.svg');
  });

  it('should return the right img for civo', () => {
    const result = providers.providerImg('civo');
    expect(result).toBe('/assets/img/providers/civo.jpeg');
  });

  it('should return the right img for kubernetes', () => {
    const result = providers.providerImg('kubernetes');
    expect(result).toBe('/assets/img/providers/kubernetes.png');
  });

  it('should return the right img for linode', () => {
    const result = providers.providerImg('linode');
    expect(result).toBe('/assets/img/providers/linode.png');
  });

  it('should return the right img for tencent', () => {
    const result = providers.providerImg('tencent');
    expect(result).toBe('/assets/img/providers/tencent.jpeg');
  });

  it('should return the right img for oci', () => {
    const result = providers.providerImg('oci');
    expect(result).toBe('/assets/img/providers/oci.png');
  });

  it('should return the right img for scaleway', () => {
    const result = providers.providerImg('scaleway');
    expect(result).toBe('/assets/img/providers/scaleway.png');
  });
});
