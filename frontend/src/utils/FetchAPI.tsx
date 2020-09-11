type ConvertedResponse = {
  converted_body: string
};

export const postKeiGo = async (kind: string, originalText: string): Promise<ConvertedResponse> => {
  let res: ConvertedResponse;
  const url = `http://34.71.216.160:3000/api/v1/keigo?kind=${kind}`;
  const body = {
    "body": originalText
  };

  // [暫定的な対応] サーバ側実装の考慮漏れを吸収するための一時的なエラー処理
  // 現状、Request Bodyの"body"に""を渡すとサーバ側でエラーが発生するため
  if (originalText === "") {
    res = {
      converted_body: originalText
    }
    return res;
  }

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
