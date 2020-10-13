type ConvertedResponse = {
  converted_body: string
};

export const postKeiGo = async (kind: string, originalText: string): Promise<ConvertedResponse> => {
  let res: ConvertedResponse;
  const url = `https://keigo-s57wlqzvfq-an.a.run.app/api/v1/keigo?kind=${kind}`;
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
    // ToDo: UI上でエラーを通知できるようにする
    console.log(error);
  }
  return res;
}
