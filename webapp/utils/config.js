import getConfig from "next/config";

const { publicRuntimeConfig } = getConfig();
const { URLS } = publicRuntimeConfig;

export const BASE_URL = URLS.BASE;
export const HUB_BASE_URL = URLS.HUB;
