import * as React from "react";
import { useState, ChangeEvent, KeyboardEvent } from "react";
import styled from "styled-components";
import Color from "../const/Color";
import { media } from "../utils/ResponsiveHelper";
import { postTranslate } from "../utils/FetchAPI"; 

type Kind = "teinei" | "sonkei" | "kenjyo";

const TranslateBox: React.SFC<{}> = props => {
  const [kind, setKind] = useState<Kind>("teinei");
  const [body, setBody] = useState<string>("");
  const [convertedBody, setConvertedBody] = useState<string>("");

  const handleRadioChange = async (event: ChangeEvent<HTMLInputElement>) => {
    const kind = event.target.id as Kind;
    setKind(kind);
    console.log(kind);
    const res = await postTranslate(kind, body);
    setConvertedBody(res.converted_body);
  };

  const handleChange = (event: ChangeEvent<HTMLInputElement>) => {
    setBody(event.target.value);
  };

  const handleKeyPress = async (event: KeyboardEvent<HTMLInputElement>) => {
    if(event.key === "Enter"){
      const res = await postTranslate(kind, body);
      setConvertedBody(res.converted_body);
    }
  };

  return (
    <Wrapper>
      <TranslateBar>
        <TranslateBarLeft>
          <LangBox>
            <label>
              原文
            </label>
          </LangBox>
        </TranslateBarLeft>
        <TranslateBarRight>
          <LangBoxRadioButton>
            <input type="radio" name="敬語" value="1" id="teinei" defaultChecked onChange={handleRadioChange} />
            <label htmlFor="teinei">丁寧</label>
          </LangBoxRadioButton>
          <LangBoxRadioButton>
            <input type="radio" name="敬語" value="2" id="sonkei" onChange={handleRadioChange} />
            <label htmlFor="sonkei">尊敬</label>
          </LangBoxRadioButton>
          <LangBoxRadioButton>
            <input type="radio" name="敬語" value="3" id="kenjyo" onChange={handleRadioChange} />
            <label htmlFor="kenjyo">謙譲</label>
          </LangBoxRadioButton>
        </TranslateBarRight>
      </TranslateBar>
      <TextBox>
        <input type="text" placeholder="テキストを入力してください" value={body} onChange={handleChange} onKeyPress={handleKeyPress} />
      </TextBox>
      <TextBox>
        <input type="text" readOnly placeholder="変換結果" value={convertedBody} />
      </TextBox>
    </Wrapper>
  );
}

const Wrapper = styled.div`
  width: calc(100% - 2 * 66px);
  padding: 66px;

  ${media.phone`
    width: 100%;
    padding: 0px;
    padding-top: 66px;
  `}
`;

const TranslateBar = styled.div`
  display: flex;
  width: 100%;
  height: 48px;

  border-top: 1px solid ${Color.BORDER};
  border-left: 1px solid ${Color.BORDER};
  border-right: 1px solid ${Color.BORDER};
`;

const TranslateBarLeft = styled.div`
  display: flex;
  width: 50%;
  height: 100%;
`;

const TranslateBarRight = styled.div`
  display: flex;
  width: 50%;
  height: 100%;
`;

const LangBox = styled.div`
  width: 150px;
  height: 48px;

  label {
    display: flex;
    align-items: center;
    justify-content: center;
    width: 100%;
    height: 100%;

    font-style: normal;
    font-weight: bold;
    font-size: 16px;
    line-height: 19px;

    color: ${Color.PRIMARY};
  }
`;

const LangBoxRadioButton = styled(LangBox)`
  input[type=radio] {
    display: none;
  }
  input[type="radio"] + label {
    color: ${Color.GRAY};
  }
  input[type="radio"]:checked + label {
    color: ${Color.PRIMARY};
    box-shadow: inset 0px -3px 0px ${Color.PRIMARY};
  }
`;

const TextBox = styled.div`
  float: left;
  width: 50%;
  height: 200px;

  input {
    width: calc(100% - 2 * 36px);
    height: calc(100% - 2 * 36px);
    outline: none;
    border: 1px solid ${Color.BORDER};

    padding: 36px;

    font-weight: 500;
    font-size: 24px;
    line-height: 28px;

    color: ${Color.BLACK};
  }

  ${media.phone`
    float: none;
    width: 100%;
  `}
`;

export default TranslateBox;
