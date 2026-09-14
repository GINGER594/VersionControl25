package main
import ("os"
        "os/exec")
func applyBashCommandAsFunctor(bashCmd, dir string, args []string) { //applies bash command as a functor to dir (other args optional)
    entities, _ := os.ReadDir(dir) //entities will be empty if err - no err-handling needed
    for _, entity := range entities {
        if entity.Name()[0] != byte('.') {
            cmd := exec.Command(bashCmd, append([]string{dir+entity.Name()}, args...)...)
            cmd.Run() //malformed commands will fail and not do anything - no err-handling needed
        }
    }
}
func main() {
    if verDir := `./.vc25/`; len(os.Args) == 3 { //defining path to version backup folder
        os.Mkdir(verDir, 0750)
        if os.Args[1] == "cpver" { //copy version (creates new if not exists)
            os.Mkdir(verDir+os.Args[2]+`/`, 0750)
            applyBashCommandAsFunctor("rm", verDir+os.Args[2]+`/`, []string{"-r"}) //removing entities from the specified version dir
            applyBashCommandAsFunctor("cp", `./`, []string{verDir+os.Args[2]+`/`, "-r"}) //copying entities from the current dir to the specified version dir
        } else if _, verExistErr := os.ReadDir(verDir+os.Args[2]+`/`); os.Args[1] == "chver" && verExistErr == nil { //change version (does nothing if not exists)
            applyBashCommandAsFunctor("rm", `./`, []string{"-r"}) //removing entities from the current dir
            applyBashCommandAsFunctor("cp", verDir+os.Args[2]+`/`, []string{`./`, "-r"}) //copying entities from the specified version dir to the current dir
        }
    }
}

//to use, run with:
//- cpver <version>
//    copies contents of current dir to <version> dir
//    creates a new version dir of name <version> if one does not exist
//
//- chver <version>
//    changes contents of current dir to contents of <version> dir
//    does nothing if a version dir of name <version> does not exist
//
//note: ignores all surface-level hidden entities
