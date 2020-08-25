/**
 * Copyright (c) 2020 TypeFox GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License-AGPL.txt in the project root for license information.
 */

import { GitHubEndpoint } from "./github-endpoint";
import { Permissions } from '../github-model/permissions';
import { GitHubRestApi } from "./github-rest-api";

export namespace TEST_TOKENS {
    export const READ_EMAIL_PERMISSION = "26def67c1bf5cb02c03abdfa1021f2ea517696d2";
    export const READ_ORG_PERMISSION = "cf0c7915e760b8a291d4bac025e00da6bc2fb600";
    export const READ_EMAIL__READ_ORG__WRITE_PUBLIC__PERMISSION = "a957b7138f205b598c532518581e42c298e9d2fc";
    export const READ_EMAIL__READ_ORG__WRITE_PRIVATE__PERMISSION = "a887d465670c4faf07a8479155b6a76cf0f106bf";
}

export class TestGitHubComEndpoint extends GitHubEndpoint {
    constructor(protected token: string) { super(); }
    async getToken(host: string, permissions?: Permissions): Promise<string> {
        return this.token;
    }
    protected get host(): string {
        return "github.com";
    }
}

export class TestGitHubComRestApi extends GitHubRestApi {
    constructor(protected token: string) { super(); }
    async getToken(host: string, permissions?: Permissions): Promise<string> {
        return this.token;
    }
    protected get host(): string {
        return "github.com";
    }
    protected get enabled(): boolean {
        return true;
    }
    protected get userAgent(): string {
        return "gitpod-test";
    }
}