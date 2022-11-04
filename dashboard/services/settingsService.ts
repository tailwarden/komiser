const BASE_URL = `http://localhost:3000`;

type Settings = {
  method: "GET" | "PUT" | "POST" | "DELETE" | undefined;
  headers: {
    "Content-Type": "application/json";
  };
  body?: string;
};

function settings(arg: string, payload?: string) {
  const settingsObj: Settings = {
    method: undefined,
    headers: {
      "Content-Type": "application/json",
    },
    body: undefined,
  };

  if (arg === "GET") {
    settingsObj.method = "GET";
    delete settingsObj.body;
  }

  if (arg === "PUT") {
    settingsObj.method = "PUT";
    delete settingsObj.body;
  }

  if (arg === "PUT" && payload) {
    settingsObj.method = "PUT";
    settingsObj.body = payload;
  }

  if (arg === "POST") {
    settingsObj.method = "POST";
    delete settingsObj.body;
  }

  if (arg === "POST" && payload) {
    settingsObj.method = "POST";
    settingsObj.body = payload;
  }

  if (arg === "DELETE") {
    settingsObj.method = "DELETE";
  }

  return settingsObj;
}

const settingsService = {
  async getInventoryStats() {
    try {
      const res = await fetch(
        `${BASE_URL}/stats`,
        settings("GET")
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async getInventoryList(urlParams: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/resources${urlParams}`,
        settings("GET")
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async searchInventory(workspaceId: string, query: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/workspaces/${workspaceId}/inventory?query=${query}`,
        settings("GET")
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },

  async saveTags(workspaceId: string, serviceId: string, payload: string) {
    try {
      const res = await fetch(
        `${BASE_URL}/workspaces/${workspaceId}/inventory/${serviceId}`,
        settings("PUT", payload)
      );
      const data = await res.json();
      return data;
    } catch (error) {
      return Error;
    }
  },
};

export default settingsService;
