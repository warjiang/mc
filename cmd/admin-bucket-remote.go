/*
 * MinIO Client (C) 2020 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package cmd

import "github.com/minio/cli"

var adminBucketRemoteCmd = cli.Command{
	Name:   "remote",
	Usage:  "manage remote bucket targets",
	Action: mainadminBucketRemote,
	Before: setGlobalsFromContext,
	Flags:  globalFlags,
	Subcommands: []cli.Command{
		adminBucketRemoteSetCmd,
		adminBucketRemoteListCmd,
		adminBucketRemoteRemoveCmd,
	},
	HideHelpCommand: true,
}

// mainadminBucketRemote is the handle for "mc admin bucket replication" command.
func mainadminBucketRemote(ctx *cli.Context) error {
	cli.ShowCommandHelp(ctx, ctx.Args().First())
	return nil
	// Sub-commands like "set", "get" have their own main.
}
