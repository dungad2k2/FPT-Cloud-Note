# Tìm hiểu Git-GitLab
## Git
Git được hiểu như là phầm mềm quản lý các local repositories. 

Trước hết để thực hiện được những công việc với GitLab, ta cần phải biết được một số câu lệnh git hay dùng:

1. `git init <tên repo>` : câu lệnh này có tác dụng tạo ra một local repo.
2. `git config --global init.defaultBranch main`: thay đổi tên default branch là main.
3. `git branch -m <tên branch>` hoặc `git branch -m <tên branch cần đổi> <tên mới>`: đổi tên branch hiện tại 
4. `git branch <tên branch>`: tạo một branch mới.
5. `git checkout -b <tên branch>`: đổi sang một branch mới.
6. `git status`: kiểm tra trạng thái của repo
7. `git add <tên file hoặc folder>`: add những file hoặc folder được thay đổi vào staging area. Lưu ý khi sử dụng câu lệnh git status những thay đổi chưa được đưa vào staging area sẽ được đánh dấu màu đỏ, ngược lại thì màu xanh.
8. `git rm --cached <file>`: để remove thay đổi khỏi staging area.
9. `git commit -m "<comment>"`: commit thay đổi vào branch.
10. `git log`: hiện log những commit trong repo.
11. `git merge <feature branch>`: thường để dùng merge branch feature vào branch main hay master.
12. `git stash push`: được hiểu như là cất giấu những thay đổi -> những thay đổi sẽ không được nhìn thấy nếu không được git stash apply hoặc git stash pop.
13. `git stash list`: liệt kê những thay đổi được cất giấu.
14. `git stash apply/pop`: đều có chức năng là hiện trở lại những thay đổi trước đó đã được cất giấu tuy nhiên apply sẽ không xóa trong `git stash list` còn pop thì có.

## GitLab CI/CD
