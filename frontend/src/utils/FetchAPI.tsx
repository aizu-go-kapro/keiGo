type ConvertedResponse = {
  converted_body: string
};

export const postKeiGo = async (kind: string, originalText: string): Promise<ConvertedResponse> => {
  let res: ConvertedResponse;
  const url = `http://34.71.216.160:3000/api/v1/keigo?kind=${kind}`;
  const body = {
    "body": originalText
  };
  try {
    const response = await fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(body)
    });
    res = await response.json();
  } catch (error) {
    console.log(error);
  }
  return res;
}
