import * as React from "react";

interface IState {
  value: string;
  convertedBody: string;
}

export default class TranslateBox extends React.Component<{}, IState> {
  constructor(props) {
    super(props);
    this.state = {
      value: "",
      convertedBody: ""
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  async handleChange(event) {
    this.setState({value: event.target.value});
    const url = "http://35.184.71.230:3000/api/v1/keigo?kind=teinei";
    const body = {
      "body": event.target.value
    };
    try {
      const response = await fetch(url, {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(body)
      });
      const data = await response.json();
      console.log(data.converted_body);
      this.setState({convertedBody: data.converted_body});
    } catch (err) {
      console.error(err);
    }
  }

  handleSubmit(event) {
    alert('An essay was submitted: ' + this.state.convertedBody);
    event.preventDefault();
  }

  render() {
    return (
      <>
        <input type="radio" name="敬語" value="丁寧" checked={true} />丁寧
        <input type="radio" name="敬語" value="尊敬" />尊敬
        <input type="radio" name="敬語" value="謙譲" />謙譲
        <form onSubmit={this.handleSubmit}>
          <label>
            原文:
            <textarea placeholder="テキストを入力してください" value={this.state.value} onChange={this.handleChange} />
          </label>
          <input type="submit" value="Submit" />
        </form>
        <textarea readOnly placeholder="変換結果" value={this.state.convertedBody} />
      </>
    );
  }
}